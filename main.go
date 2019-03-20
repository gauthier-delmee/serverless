package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Welcome to 'Hands-On serverless Applications with Go'",
	}, nil
}

func main() {
	lambda.Start(handler)
}
