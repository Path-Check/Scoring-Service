package model

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveJSONFile(t *testing.T) {
	requestData := []byte(`
        {
            "newExposureSummary":
            {
                "dateReceived": 1597482000,
                "timeZoneOffset": 32400,
                "seqNoInDay": 1,
                "attenuationDurations": {"low": 1800, "medium": 1800, "high": 0},
                "matchedKeyCount": 4,
                "daysSinceLastExposure": 1,
                "maximumRiskScore": 1,
                "riskScoreSum": 1
            },
            "unusedExposureSummaries":
            [{
                "dateReceived": 1597395600,
                "timeZoneOffset": 32400,
                "seqNoInDay": 1,
                "attenuationDurations": {"low": 600, "medium": 0, "high": 0},
                "matchedKeyCount": 2,
                "daysSinceLastExposure": 0,
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
	err := json.Unmarshal(requestData, &parsedRequest)
	if err != nil {
		log.Println(err)
	}

	logResponse, err := SaveJSONFile(&parsedRequest)
	if err != nil {
		log.Println(err)
	}

	assert.Equal(t, "Saved ExposureNotificationRequest in JSON file", logResponse)
}
