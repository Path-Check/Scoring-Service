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
            "newExposureSummary":
            {
                "dateReceived": 1597482000,
                "timezoneOffset": 32400,
                "seqNoInDay": 1,
                "attenuationDurations": {"low": 180, "medium": 60, "high": 500},
                "matchedKeyCount": 1,
                "daysSinceLastExposure": 1,
                "maximumRiskScore": 1,
                "riskScoreSum": 1
            },
            "unusedExposureSummaries":
            [{
                "dateReceived": 1597482000,
                "timezoneOffset": 32400,
                "seqNoInDay": 2,
                "attenuationDurations": {"low": 0, "medium": 300, "high": 300},
                "matchedKeyCount": 1,
                "daysSinceLastExposure": 3,
                "maximumRiskScore": 1,
                "riskScoreSum": 1
			}],
            "exposureConfiguration":
            {
                "minimumRiskScore": 0,
                "attenuationDurationThresholds": [53, 60],
                "attenuationLevelValues": [1,2,3,4,5,6,7,8],
                "daysSinceLastExposureLevelValues": [1,2,3,4,5,6,7,8],
                "durationLevelValues": [1,2,3,4,5,6,7,8],
                "transmissionRiskLevelValues": [1,2,3,4,5,6,7,8],
                "attenuationBucketWeights": [1, 0.5, 0],
                "triggerThresholdWeightedDuration": 15
            }
        }`)

	var parsedRequest ExposureNotificationRequest
	error := json.Unmarshal(requestData, &parsedRequest)
	if error != nil {
		log.Println(error)
	}

	assert.Equal(t, 1597482000, parsedRequest.UnusedExposureSummaries[0].DateReceived, "dateReceived did not parse correctly")
	assert.Equal(t, []float32{1, 0.5, 0}, parsedRequest.ExposureConfiguration.AttenuationBucketWeights,
		"attenuationBucketWeights did not parse correctly.")
}

func TestWriteResponse(t *testing.T) {
	responseData := &ExposureNotificationResponse{
		Notifications: []Notification{
			{
				ExposureSummaries: []ExposureSummary{
					{
						DateReceived:          1597654800,
						TimezoneOffset:        0,
						SeqNoInDay:            2,
						AttenuationDurations:  AttenuationDurations{Low: 1800, Medium: 0, High: 0},
						MatchedKeyCount:       1,
						DaysSinceLastExposure: 0,
						MaximumRiskScore:      1,
						RiskScoreSum:          1,
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

	expected :=
		`{"notifications":[{"exposureSummaries":[{"dateReceived":1597654800,"timezoneOffset":0,"seqNoInDay":2,"attenuationDurations":{"low":1800,"medium":0,"high":0},"matchedKeyCount":1,"daysSinceLastExposure":0,"maximumRiskScore":1,"riskScoreSum":1}],"durationSeconds":1800,"dateOfExposure":1597482000}]}`

	assert.Equal(t, expected, string(response))
}
