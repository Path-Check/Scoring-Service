package model

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRequest(t *testing.T) {
	requestData := []byte(`
        {
            "new_exposure_summary":
            {
                "date_received": 1597482000,
                "timezone_offset": 32400,
                "seq_no_in_day": 1,
                "attenuation_durations": {"low": 180, "medium": 60, "high": 500},
                "matched_key_count": 1,
                "days_since_last_exposure": 1,
                "maximum_risk_score": 1,
                "risk_score_sum": 1
            },
            "unused_exposure_summaries":
            [{
                "date_received": 1597482000,
                "timezone_offset": 32400,
                "seq_no_in_day": 2,
                "attenuation_durations": {"low": 0, "medium": 300, "high": 300},
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

	assert.Equal(t, 1597482000, parsedRequest.UnusedExposureSummaries[0].DateReceived,
		"date_received did not parse correctly.")
}

func TestWriteResponse(t *testing.T) {
	responseData := &ExposureNotificationResponse{
		Notifications: []Notification{
			{
				ExposureSummaries: []ExposureSummary{
					{
						DateReceived:         1597654800,
						SeqNoInDay:           2,
						AttenuationDurations: AttenuationDurations{Low: 1800, Medium: 0, High: 0},
						MatchedKeyCount:      1,
					},
				},
				DurationSeconds: 1800,
				DateOfExposure:  1597482000,
			},
		},
	}
	response, error := json.Marshal(responseData)
	if error != nil {
		log.Println(error)
	}

	expected := `{"notifications":[{"exposure_summaries":[{"date_received":1597654800,"seq_no_in_day":2,"attenuation_durations":{"low":1800,"medium":0,"high":0},"matched_key_count":1}],"duration_seconds":1800,"date_of_exposure":1597482000}]}`

	assert.Equal(t, expected, string(response))
}
