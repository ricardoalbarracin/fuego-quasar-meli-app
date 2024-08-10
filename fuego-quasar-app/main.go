package main

import (
	"fuego-quasar-app/internal/infrastructure/di"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambdaHandler := di.InitializeMyService()
	lambda.Start(lambdaHandler.HandleRequest)
}
