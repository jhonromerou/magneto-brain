package awscommon

import (
	"testing"

	"github.com/jhonromerou/magneto-brain/src/domain"
)

func Test_Aws_Dynamodb_Provider_ConfigProvider(t *testing.T) {

	tests := []struct {
		name      string
		wantError bool
	}{
		{
			name:      "error creating provider",
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ConfigProvider()

			if err != nil != tt.wantError {
				t.Fatal(domain.TestingErrorNotMatched(1, err, tt.wantError))
			}
		})
	}
}
