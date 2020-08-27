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

	assert.Equal(t, expected, string(response))
}

func TestAggregatedExposuresDeterministicDayDifferentDays(t *testing.T) {
	// In this case, there is not sufficient exposure in the new ExposureSummary
	// for a notification. The second exposure occurred on a different day,
	// so they should not be aggregated, and there will be no notification.
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
                "days_since_last_exposure": 3,
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

	expected := `{}`
	assert.Equal(t, expected, string(response))
}

func TestAggregatedExposuresNonDeterministicDay(t *testing.T) {
	// In this case, there is not sufficient exposure in the new ExposureSummary
	// for a notification. The second exposure came from multiple keys, so
	// we don't want to use that for aggregation since we can't tell how
	// much exposure there were per day.
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
                "matched_key_count": 2,
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

	expected := `{}`
	assert.Equal(t, expected, string(response))
}

func TestNonDeterministicDayAvgAboveThreshold(t *testing.T) {
	// If the new exposure matches 2-3 keys, but the average is still above
	// the threshold of 15 minutes, we should still notify since we can tell
	// that at least one exposure was above 15 minutes long.
	requestData := []byte(`
        {
            "new_exposure_summary":
            {
                "date_received": 1597482000,
                "timezone_offset": 32400,
                "seq_no_in_day": 1,
                "attenuation_durations": {"low": 1900, "medium": 0, "high": 0},
                "matched_key_count": 2,
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

	expected := `{"notifications":[{"exposure_summaries":[{"date_received":1597482000,"timezone_offset":32400,"seq_no_in_day":1,"attenuation_durations":{"low":1900,"medium":0,"high":0},"matched_key_count":2,"days_since_last_exposure":1,"maximum_risk_score":1,"risk_score_sum":1}],"duration_seconds":1900,"date_most_recent_exposure":1597395600,"matched_key_count":2}]}`

	assert.Equal(t, expected, string(response))
}

func TestNoExposureError(t *testing.T) {
	requestData := []byte(`
        {
            "new_exposure_summary":
            {
                "date_received": 1597482000,
                "timezone_offset": 32400,
                "seq_no_in_day": 1,
                "attenuation_durations": {"low": 0, "medium": 0, "high": 0},
                "matched_key_count": 0,
                "days_since_last_exposure": 0,
                "maximum_risk_score": 0,
                "risk_score_sum": 0
            }
        }`)

	var parsedRequest ExposureNotificationRequest
	error := json.Unmarshal(requestData, &parsedRequest)
	if error != nil {
		log.Println(error)
	}

	responseData, scoreError := ScoreV1(&parsedRequest)

	response, error := json.Marshal(responseData)
	if error != nil {
		log.Println(error)
	}
	assert.Equal(t, "{}", string(response))
	assert.Equal(t, "Matched key count was 0.", scoreError.Error())
}

