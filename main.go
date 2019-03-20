package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

func handler() (string, error) {
	return "Welcome to 'Hands-On serverless Applications with Go'", nil
}

func main() {
	lambda.Start(handler)
}
