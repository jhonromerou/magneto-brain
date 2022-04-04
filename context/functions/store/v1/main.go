package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jhonromerou/magneto-brain/context/functions/store/v1/internal/di"
)

func main() {
	handler, _ := di.Initialize()
	lambda.Start(handler.Handle)
}
