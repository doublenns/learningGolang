package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// type Response struct {
// 	StatusCode int    `json:"statusCode"`
// 	Body       string `json:"body"`
// }

// func handler() (Response, error) {
// 	return Response{
// 		StatusCode: 200,
// 		Body:       "Welcome to Serverless World",
// 	}, nil
// }

func handler() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Welcome to Serverless World",
	}, nil
}

func main() {
	lambda.Start(handler)
}
