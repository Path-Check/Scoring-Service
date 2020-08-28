package handler

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenericHandler(t *testing.T) {
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

	// use RawMessage instead of Marshal to avoid escaping characters
	rm := json.RawMessage(requestData)

	rawJSON, _ := rm.MarshalJSON()
	body := string(rawJSON)

	expectedResponse := "{\"notifications\":[{\"exposureSummaries\":[{\"dateReceived\":1597482000,\"timezoneOffset\":32400,\"seqNoInDay\":1,\"attenuationDurations\":{\"low\":900,\"medium\":0,\"high\":0},\"matchedKeyCount\":1,\"daysSinceLastExposure\":1,\"maximumRiskScore\":1,\"riskScoreSum\":1}],\"durationSeconds\":900,\"dateOfExposure\":1597395600}]}"

	// Test a successful request.
	status, response, err := GenericHandler(body)
	assert.Equal(t, 200, status)
	assert.Equal(t, nil, err)
	assert.Equal(t, expectedResponse, response)
}

func TestEmptyBodyError(t *testing.T) {
	// Test that we get an error on an empty body.
	status, response, err := GenericHandler("")
	assert.Equal(t, 400, status)
	assert.Equal(t, ErrRequestBodyEmpty, err)
	assert.Equal(t, "", response)
}

func TestMalformedRequestError(t *testing.T) {
	status, response, err := GenericHandler("i am not a valid JSON string")
	assert.Equal(t, 400, status)
	assert.Equal(t, ErrUnmarshalJSON, err)
	assert.Equal(t, "", response)
}
