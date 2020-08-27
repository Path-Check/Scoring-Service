package handler

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlerSuccess(t *testing.T) {
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

func TestMalformedRequstError(t *testing.T) {
	status, response, err := GenericHandler("i am not a valid JSON string")
	assert.Equal(t, 400, status)
	assert.Equal(t, ErrUnmarshalJSON, err)
	assert.Equal(t, "", response)
}
