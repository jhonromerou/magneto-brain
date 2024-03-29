// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/jhonromerou/magneto-brain/context/functions/stats/v1/internal"
	"github.com/jhonromerou/magneto-brain/src/infrastructure/awscommon"
	"github.com/jhonromerou/magneto-brain/src/infrastructure/awsdynamodb"
	"github.com/jhonromerou/magneto-brain/src/infrastructure/dotenv"
	"github.com/jhonromerou/magneto-brain/src/infrastructure/logruslogger"
	"github.com/jhonromerou/magneto-brain/src/repositories"
	"github.com/sirupsen/logrus"
)

// Injectors from wire.go:

func Initialize() (*internal.Handler, error) {
	logger := logruslogger.NewLogrusLoggerProvider()
	entry := logrus.NewEntry(logger)
	logrusLogger := logruslogger.NewLogrusLogger(entry)
	config, err := awscommon.ConfigProvider()
	if err != nil {
		return nil, err
	}
	v := awsdynamodb.DynamodbOptions()
	client := dynamodb.NewFromConfig(config, v...)
	awsDynamoDb := awsdynamodb.NewAwsDynamoDb(client)
	dotEnvEnvironmentReposity := dotenv.NewDotEnvEnvironmentReposity()
	dynamodbStatsRepository := repositories.NewDynamodbStatsRepository(awsDynamoDb, dotEnvEnvironmentReposity)
	handler := internal.NewHandler(logrusLogger, dynamodbStatsRepository)
	return handler, nil
}
