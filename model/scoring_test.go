package model

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOneExposure(t *testing.T) {
	requestData := []byte(`
        {
            "new_exposure_summary":
            {
                "date_received": 1597482000,
                "timezone_offset": 32400,
                "seq_no_in_day": 1,
                "attenuation_durations": {"low": 900, "medium": 0, "high": 0},
                "matched_key_count": 1,
                "days_since_last_exposure": 1,
                "maximum_risk_score": 1,
                "risk_score_sum": 1
            }
        }`)

	var parsedRequest ExposureNotificationRequest
	error := json.Unmarshal(requestData, &parsedRequest)
	if error != nil {
		log.Println(error)
	}

	responseData, _ := ScoreV1(&parsedRequest)

	response, error := json.Marshal(responseData)
	if error != nil {
		log.Println(error)
	}

	expected := `{"notifications":[{"exposure_summaries":[{"date_received":1597482000,"timezone_offset":32400,"seq_no_in_day":1,"attenuation_durations":{"low":900,"medium":0,"high":0},"matched_key_count":1,"days_since_last_exposure":1,"maximum_risk_score":1,"risk_score_sum":1}],"duration_seconds":900,"date_of_exposure":1597395600}]}`

	assert.Equal(t, expected, string(response))
}

func TestInsufficientExposure(t *testing.T) {
	requestData := []byte(`
        {
            "new_exposure_summary":
            {
                "date_received": 1597482000,
                "timezone_offset": 32400,
                "seq_no_in_day": 1,
                "attenuation_durations": {"low": 90, "medium": 0, "high": 0},
                "matched_key_count": 1,
                "days_since_last_exposure": 1,
                "maximum_risk_score": 1,
                "risk_score_sum": 1
            }
        }`)

	var parsedRequest ExposureNotificationRequest
	error := json.Unmarshal(requestData, &parsedRequest)
	if error != nil {
		log.Println(error)
	}

	responseData, _ := ScoreV1(&parsedRequest)

	response, error := json.Marshal(responseData)
	if error != nil {
		log.Println(error)
	}

	expected := `{}`

	assert.Equal(t, expected, string(response))
}

func TestAggregatedExposuresDeterministicDay(t *testing.T) {
	// In this case, there is not sufficient exposure in the new ExposureSummary
	// for a notification. But in aggregation with an older summary, there is.
	requestData := []byte(`
        {
            "new_exposure_summary":
            {
                "date_received": 1597482000,
                "timezone_offset": 32400,
                "seq_no_in_day": 1,
                "attenuation_durations": {"low": 400, "medium": 0, "high": 0},
                "matched_key_count": 1,
                "days_since_last_exposure": 1,
                "maximum_risk_score": 1,
                "risk_score_sum": 1
            },
            "unused_exposure_summaries":
            [{
                "date_received": 1597395600,
                "timezone_offset": 32400,
                "seq_no_in_day": 1,
                "attenuation_durations": {"low": 600, "medium": 0, "high": 0},
                "matched_key_count": 1,
                "days_since_last_exposure": 0,
                "maximum_risk_score": 1,
                "risk_score_sum": 1
            }]
        }`)

	var parsedRequest ExposureNotificationRequest
	error := json.Unmarshal(requestData, &parsedRequest)
	if error != nil {
		log.Println(error)
	}

	responseData, _ := ScoreV1(&parsedRequest)
	response, error := json.Marshal(responseData)
	if error != nil {
		log.Println(error)
	}

	expected := `{"notifications":[{"exposure_summaries":[{"date_received":1597482000,"timezone_offset":32400,"seq_no_in_day":1,"attenuation_durations":{"low":400,"medium":0,"high":0},"matched_key_count":1,"days_since_last_exposure":1,"maximum_risk_score":1,"risk_score_sum":1},{"date_received":1597395600,"timezone_offset":32400,"seq_no_in_day":1,"attenuation_durations":{"low":600,"medium":0,"high":0},"matched_key_count":1,"days_since_last_exposure":0,"maximum_risk_score":1,"risk_score_sum":1}],"duration_seconds":1000,"date_of_exposure":1597395600}]}`

	// SUPER IMPORTANT TODO!!!!!! This one should be in the JSON string for the
	// second summary above, but thanks to omitempty and 0 being the default value,
	// it just ends up being left out. We need to omit fields that were not set, but
	// still use them if they have a value set that also happens to be the default
	// value. Fix this.
	//          "days_since_last_exposure": 0,

	assert.Equal(t, expected, string(response))
}
