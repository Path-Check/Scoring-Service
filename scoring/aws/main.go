package main

import (
	"errors"
	"log"

	"github.com/Path-Check/Scoring-Service/scoring/common"
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
	// Call the generic handler common to all Clouds.
	statusCode, responseString, err := handler.GenericHandler(request.Body)
	if err != nil {
		// stdout and stderr are sent to AWS CloudWatch Logs
		log.Printf("Error: %s", err.Error())
        }

	return events.APIGatewayProxyResponse{
		Body:       responseString,
		StatusCode: statusCode,
	}, err
}

func main() {
	lambda.Start(Handler)
}
