package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

// Response defines the structure complying to Lambda function requirements
type Response struct {
	StatusCode int    `json:"statusCode"`
	Body       string `json:"body"`
}

func handler() (Response, error) {
	return Response{
		StatusCode: 200,
		Body:       "Welcome to 'Hands-On serverless Applications with Go'",
	}, nil
}

func main() {
	lambda.Start(handler)
}
