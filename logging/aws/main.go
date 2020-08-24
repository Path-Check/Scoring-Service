package main

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	ErrRequestBodyEmpty = errors.New("HTTP request body is empty")
)

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// You may also use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request ID: %s\n", request.RequestContext.RequestID)

	// If HTTP request body is empty, throw error
	if len(request.Body) < 1 {
		return events.APIGatewayProxyResponse{}, ErrRequestBodyEmpty
	}
	log.Println("API Gateway request body:")
	log.Println(request.Body)

	enReq := LogRequest{}
	err := json.Unmarshal([]byte(request.Body), &enReq)
	if err != nil {
		log.Println("Can't unmarshal APIGatewayProxyRequest body")
	}
	// What should we log?
	enRes, err := Logger(&enReq)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
	}

	response, err := json.Marshal(&enRes)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
	}
	log.Println("Scoring response:")
	log.Println(string(response))

	return events.APIGatewayProxyResponse{
		Body:       string(response),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
