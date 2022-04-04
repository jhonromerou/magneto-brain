package awssqs

import (
	"errors"
	"testing"

	"github.com/jhonromerou/magneto-brain/src/domain"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/stretchr/testify/mock"
)

func Test_Aws_Sqs(t *testing.T) {
	type fields struct {
		awsSqsClient *MockAwsSqsClient
	}
	tests := []struct {
		name      string
		args      interface{}
		fields    fields
		funcMock  func(fields)
		wantError bool
		want      interface{}
	}{
		{
			name: "error geting url",
			fields: fields{
				awsSqsClient: &MockAwsSqsClient{},
			},
			funcMock: func(f fields) {
				queueNameInput := &sqs.GetQueueUrlInput{
					QueueName: aws.String("queue-name"),
				}

				queueOutput := &sqs.GetQueueUrlOutput{}
				err := errors.New("some error")
				f.awsSqsClient.On("GetQueueUrl", mock.Anything, queueNameInput).Once().Return(queueOutput, err)
			},
			wantError: true,
		},
		{
			name: "error publishing message",
			fields: fields{
				awsSqsClient: &MockAwsSqsClient{},
			},
			funcMock: func(f fields) {
				queueName := aws.String("queue-name")
				queueNameInput := &sqs.GetQueueUrlInput{
					QueueName: queueName,
				}
				queueOutput := &sqs.GetQueueUrlOutput{}
				params := &sqs.SendMessageInput{
					MessageBody: aws.String("message"),
					QueueUrl:    queueOutput.QueueUrl,
				}
				err := errors.New("some error")
				f.awsSqsClient.On("GetQueueUrl", mock.Anything, queueNameInput).Once().Return(queueOutput, nil)
				f.awsSqsClient.On("SendMessage", mock.Anything, params).Once().Return(nil, err)
			},
			wantError: true,
		},
		{
			name: "success publishing message",
			fields: fields{
				awsSqsClient: &MockAwsSqsClient{},
			},
			funcMock: func(f fields) {
				queueName := aws.String("queue-name")
				queueNameInput := &sqs.GetQueueUrlInput{
					QueueName: queueName,
				}
				queueOutput := &sqs.GetQueueUrlOutput{}
				params := &sqs.SendMessageInput{
					MessageBody: aws.String("message"),
					QueueUrl:    queueOutput.QueueUrl,
				}
				f.awsSqsClient.On("GetQueueUrl", mock.Anything, queueNameInput).Once().Return(queueOutput, nil)
				f.awsSqsClient.On("SendMessage", mock.Anything, params).Once().Return(nil, nil)
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.funcMock(tt.fields)
			awsSqs := NewAwsSqs(tt.fields.awsSqsClient)
			awsSqs.SetQueueName("queue-name")
			err := awsSqs.PublishMessage("message")
			if (err != nil) != tt.wantError {
				t.Fatal(domain.TestingErrorNotMatched(1, err, tt.wantError))
			}
			tt.fields.awsSqsClient.AssertExpectations(t)
		})
	}
}
