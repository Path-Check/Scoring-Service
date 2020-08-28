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
            "newExposureSummary":
            {
                "dateReceived": 1597482000,
                "timezoneOffset": 32400,
                "seqNoInDay": 1,
                "attenuationDurations": {"low": 900, "medium": 0, "high": 0},
                "matchedKeyCount": 1,
                "daysSinceLastExposure": 1,
                "maximumRiskScore": 1,
                "riskScoreSum": 1
            },
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

	responseData, _ := ScoreV1(&parsedRequest)

	response, error := json.Marshal(responseData)
	if error != nil {
		log.Println(error)
	}

	expected := `{"notifications":[{"exposureSummaries":[{"dateReceived":1597482000,"timezoneOffset":32400,"seqNoInDay":1,"attenuationDurations":{"low":900,"medium":0,"high":0},"matchedKeyCount":1,"daysSinceLastExposure":1,"maximumRiskScore":1,"riskScoreSum":1}],"durationSeconds":900,"dateOfExposure":1597395600}]}`

	assert.Equal(t, expected, string(response))
}

func TestInsufficientExposure(t *testing.T) {
	requestData := []byte(`
        {
            "newExposureSummary":
            {
                "dateReceived": 1597482000,
                "timezoneOffset": 32400,
                "seqNoInDay": 1,
                "attenuationDurations": {"low": 90, "medium": 0, "high": 0},
                "matchedKeyCount": 1,
                "daysSinceLastExposure": 1,
                "maximumRiskScore": 1,
                "riskScoreSum": 1
            },
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
            "newExposureSummary":
            {
                "dateReceived": 1597482000,
                "timezoneOffset": 32400,
                "seqNoInDay": 1,
                "attenuationDurations": {"low": 400, "medium": 0, "high": 0},
                "matchedKeyCount": 1,
                "daysSinceLastExposure": 1,
                "maximumRiskScore": 1,
                "riskScoreSum": 1
            },
            "unusedExposureSummaries":
            [{
                "dateReceived": 1597395600,
                "timezoneOffset": 32400,
                "seqNoInDay": 1,
                "attenuationDurations": {"low": 600, "medium": 0, "high": 0},
                "matchedKeyCount": 1,
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
	error := json.Unmarshal(requestData, &parsedRequest)
	if error != nil {
		log.Println(error)
	}

	responseData, _ := ScoreV1(&parsedRequest)
	response, error := json.Marshal(responseData)
	if error != nil {
		log.Println(error)
	}

	expected := `{"notifications":[{"exposureSummaries":[{"dateReceived":1597482000,"timezoneOffset":32400,"seqNoInDay":1,"attenuationDurations":{"low":400,"medium":0,"high":0},"matchedKeyCount":1,"daysSinceLastExposure":1,"maximumRiskScore":1,"riskScoreSum":1},{"dateReceived":1597395600,"timezoneOffset":32400,"seqNoInDay":1,"attenuationDurations":{"low":600,"medium":0,"high":0},"matchedKeyCount":1,"daysSinceLastExposure":0,"maximumRiskScore":1,"riskScoreSum":1}],"durationSeconds":1000,"dateOfExposure":1597395600}]}`

	assert.Equal(t, expected, string(response))
}

func TestAggregatedExposuresDeterministicDayDifferentDays(t *testing.T) {
	// In this case, there is not sufficient exposure in the new ExposureSummary
	// for a notification. The second exposure occurred on a different day,
	// so they should not be aggregated, and there will be no notification.
	requestData := []byte(`
        {
            "newExposureSummary":
            {
                "dateReceived": 1597482000,
                "timezoneOffset": 32400,
                "seqNoInDay": 1,
                "attenuationDurations": {"low": 400, "medium": 0, "high": 0},
                "matchedKeyCount": 1,
                "daysSinceLastExposure": 1,
                "maximumRiskScore": 1,
                "riskScoreSum": 1
            },
            "unusedExposureSummaries":
            [{
                "dateReceived": 1597395600,
                "timezoneOffset": 32400,
                "seqNoInDay": 1,
                "attenuationDurations": {"low": 600, "medium": 0, "high": 0},
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
            "newExposureSummary":
            {
                "dateReceived": 1597482000,
                "timezoneOffset": 32400,
                "seqNoInDay": 1,
                "attenuationDurations": {"low": 400, "medium": 0, "high": 0},
                "matchedKeyCount": 1,
                "daysSinceLastExposure": 1,
                "maximumRiskScore": 1,
                "riskScoreSum": 1
            },
            "unusedExposureSummaries":
            [{
                "dateReceived": 1597395600,
                "timezoneOffset": 32400,
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
            "newExposureSummary":
            {
                "dateReceived": 1597482000,
                "timezoneOffset": 32400,
                "seqNoInDay": 1,
                "attenuationDurations": {"low": 1900, "medium": 0, "high": 0},
                "matchedKeyCount": 2,
                "daysSinceLastExposure": 1,
                "maximumRiskScore": 1,
                "riskScoreSum": 1
            },
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

	responseData, _ := ScoreV1(&parsedRequest)
	response, error := json.Marshal(responseData)
	if error != nil {
		log.Println(error)
	}

	expected := `{"notifications":[{"exposureSummaries":[{"dateReceived":1597482000,"timezoneOffset":32400,"seqNoInDay":1,"attenuationDurations":{"low":1900,"medium":0,"high":0},"matchedKeyCount":2,"daysSinceLastExposure":1,"maximumRiskScore":1,"riskScoreSum":1}],"durationSeconds":1900,"dateMostRecentExposure":1597395600,"matchedKeyCount":2}]}`

	assert.Equal(t, expected, string(response))
}

func TestNoExposureError(t *testing.T) {
	requestData := []byte(`
        {
            "newExposureSummary":
            {
                "dateReceived": 1597482000,
                "timezoneOffset": 32400,
                "seqNoInDay": 1,
                "attenuationDurations": {"low": 0, "medium": 0, "high": 0},
                "matchedKeyCount": 0,
                "daysSinceLastExposure": 0,
                "maximumRiskScore": 0,
                "riskScoreSum": 0
            },
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

	responseData, scoreError := ScoreV1(&parsedRequest)

	response, error := json.Marshal(responseData)
	if error != nil {
		log.Println(error)
	}
	assert.Equal(t, "{}", string(response))
	assert.Equal(t, "Matched key count was 0.", scoreError.Error())
}
