package handler

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlerSuccess(t *testing.T) {
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

	// use RawMessage instead of Marshal to avoid escaping characters
	rm := json.RawMessage(requestData)

	rawJSON, _ := rm.MarshalJSON()
	body := string(rawJSON)

	expectedResponse := "{\"notifications\":[{\"exposure_summaries\":[{\"date_received\":1597482000,\"timezone_offset\":32400,\"seq_no_in_day\":1,\"attenuation_durations\":{\"low\":900,\"medium\":0,\"high\":0},\"matched_key_count\":1,\"days_since_last_exposure\":1,\"maximum_risk_score\":1,\"risk_score_sum\":1}],\"duration_seconds\":900,\"date_of_exposure\":1597395600}]}"

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
