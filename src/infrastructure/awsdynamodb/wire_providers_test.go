package awsdynamodb

import (
	"reflect"
	"testing"

	"github.com/jhonromerou/magneto-brain/src/domain"
	"github.com/jhonromerou/magneto-brain/src/infrastructure/awscommon"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func Test_Aws_Dynamodb_Provider_DynamodbConfigProvider(t *testing.T) {
	tests := []struct {
		name string
		want *awscommon.ConfigOptions
	}{
		{
			name: "success create provider",
			want: &awscommon.ConfigOptions{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := DynamodbConfigProvider()
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatal(domain.TestingErrorNotMatched(1, got, tt.want))
			}
		})
	}
}

func Test_Aws_Dynamodb_Provider_DynamodbOptions(t *testing.T) {
	tests := []struct {
		name string
		want []func(*dynamodb.Options)
	}{
		{
			name: "success create provider",
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DynamodbOptions()
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatal(domain.TestingErrorNotMatched(1, got, tt.want))
			}
		})
	}
}
