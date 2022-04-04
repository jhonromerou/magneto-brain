package awssqs

import (
	"github.com/jhonromerou/magneto-brain/src/infrastructure/awscommon"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/google/wire"
)

// Set groups dependencies for the creation of aws services.
var SetAwsSqs = wire.NewSet(
	SqsConfigProvider,
	SqsOptions,
	awscommon.ConfigProvider,
	sqs.NewFromConfig,
	NewAwsSqs,
	wire.Bind(new(AwsSqsClient), new(*sqs.Client)),
)

// SqsConfigProvider sets the default configuration of aws sqs.
func SqsConfigProvider() (*awscommon.ConfigOptions, error) {
	return &awscommon.ConfigOptions{}, nil
}

// SqsOptions sets the default configuration of aws sqs.
func SqsOptions() []func(*sqs.Options) {
	return nil
}
