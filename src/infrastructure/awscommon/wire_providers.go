package awscommon

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

// ConfigOptions defines aws configuration options
type ConfigOptions struct {
	Region   string
	Endpoint string
}

// ConfigProvider loads the base session for aws service clients
func ConfigProvider() (aws.Config, error) {
	var opts config.LoadOptionsFunc = func(*config.LoadOptions) error {
		return nil
	}

	ctx := context.TODO()

	awsCfg, err := config.LoadDefaultConfig(ctx, opts)
	if err != nil {
		return aws.Config{}, errors.New("cannot load the aws config")
	}

	return awsCfg, nil
}
