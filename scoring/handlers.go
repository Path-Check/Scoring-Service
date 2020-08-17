package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	enReq := ExposureNotificationRequest{}
	err := json.Unmarshal([]byte(req.Body), &enReq)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
	}
	enRes := ExposureNotificationResponse{}
	response, err := json.Marshal(&enRes)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
	}
	return events.APIGatewayProxyResponse{Body: string(response), StatusCode: 200}, nil
}
