package awsdynamodb

import (
	"github.com/jhonromerou/magneto-brain/src/domain"
	"github.com/jhonromerou/magneto-brain/src/infrastructure/awscommon"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/wire"
)

var SetAwsDynamodb = wire.NewSet(
	DynamodbOptions,
	DynamodbConfigProvider,
	awscommon.ConfigProvider,
	dynamodb.NewFromConfig,
	NewAwsDynamoDb,
	wire.Bind(new(AwsDynamoDbClient), new(*dynamodb.Client)),
	wire.Bind(new(domain.DatabaseRepository), new(*AwsDynamoDb)),
)

// DynamodbConfigProvider sets the default configuration of aws dynamodb
func DynamodbConfigProvider() (*awscommon.ConfigOptions, error) {
	return &awscommon.ConfigOptions{}, nil
}

// DynamodbOptions sets the default configuration of aws dynamodb.
func DynamodbOptions() []func(*dynamodb.Options) {
	return nil
}
