package main

import (
	"github.com/jhonromerou/magneto-brain/context/functions/stats/v1/internal/di"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	handler, _ := di.Initialize()
	lambda.Start(handler.Handle)
}
