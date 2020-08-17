package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	scoreRequest := ScoreRequest{}
	err := json.Unmarshal([]byte(req.Body), &scoreRequest)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
	}
	scoreResponse := ScoreResponse{}
	response, err := json.Marshal(&scoreResponse)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
	}
	return events.APIGatewayProxyResponse{Body: string(response), StatusCode: 200}, nil
}
