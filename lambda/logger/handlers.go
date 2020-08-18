package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logRequest := LogRequest{}
	err := json.Unmarshal([]byte(req.Body), &logRequest)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
	}
	logResponse := LogResponse{}
	response, err := json.Marshal(&logResponse)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
	}
	return events.APIGatewayProxyResponse{Body: string(response), StatusCode: 200}, nil
}
