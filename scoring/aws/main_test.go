package main

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
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
	rawData := string(rawJSON)

	apigRequest := events.APIGatewayProxyRequest{
		HTTPMethod: "PUT",
		Path:       "/v1/score",
		Body:       rawData,
		Headers: map[string]string{
			"Content-Type": "application/json",
			"X-Foo":        "bar",
			"Host":         "example.com",
		},
		RequestContext: events.APIGatewayProxyRequestContext{
			RequestID: "1234",
			Stage:     "prod",
		},
	}

	tests := []struct {
		request events.APIGatewayProxyRequest
		expect  string
		err     error
	}{
		{
			// Test scoring lambda function responds right
			// when exposure summary data put in API Gateway request
			request: apigRequest,
			expect:  "{\"notifications\":[{\"exposureSummaries\":[{\"dateReceived\":1597482000,\"timezoneOffset\":32400,\"seqNoInDay\":1,\"attenuationDurations\":{\"low\":900,\"medium\":0,\"high\":0},\"matchedKeyCount\":1,\"daysSinceLastExposure\":1,\"maximumRiskScore\":1,\"riskScoreSum\":1}],\"durationSeconds\":900,\"dateOfExposure\":1597395600}]}",
			err:     nil,
		},
		{
			// Test lambda function responds ErrRequestBodyEmpty
			// when API Gateway request body is empty
			request: events.APIGatewayProxyRequest{Body: ""},
			expect:  "",
			err:     ErrRequestBodyEmpty,
		},
	}

	for _, test := range tests {
		response, err := Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response.Body)
	}
}
