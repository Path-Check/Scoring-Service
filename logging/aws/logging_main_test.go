package main

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestLoggingHandler(t *testing.T) {
	requestData := []byte(`
	{
		"newExposureSummary":
		{
			"dateReceived": 1597482000,
			"timeZoneOffset": 32400,
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
		HTTPMethod: "POST",
		Path:       "/v1/log",
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
			request: apigRequest,
			expect:  "\"Saved ExposureNotificationRequest in JSON file\"",
			err:     nil,
		},
		{
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
