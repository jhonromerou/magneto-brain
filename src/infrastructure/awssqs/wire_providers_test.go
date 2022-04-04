package awssqs

import (
	"reflect"
	"testing"

	"github.com/jhonromerou/magneto-brain/src/domain"
	"github.com/jhonromerou/magneto-brain/src/infrastructure/awscommon"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func Test_Aws_Dynamodb_Provider_SqsConfigProvider(t *testing.T) {
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
			got, _ := SqsConfigProvider()
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatal(domain.TestingErrorNotMatched(1, got, tt.want))
			}
		})
	}
}

func Test_Aws_Dynamodb_Provider_SqsOptions(t *testing.T) {
	tests := []struct {
		name string
		want []func(*sqs.Options)
	}{
		{
			name: "success create provider",
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SqsOptions()
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatal(domain.TestingErrorNotMatched(1, got, tt.want))
			}
		})
	}
}
