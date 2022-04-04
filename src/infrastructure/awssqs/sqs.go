package awssqs

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

const connectionTimeout = 25 * time.Second

type AwsSqsClient interface {
	SendMessage(ctx context.Context, params *sqs.SendMessageInput, optFns ...func(*sqs.Options)) (*sqs.SendMessageOutput, error)
	GetQueueUrl(ctx context.Context, params *sqs.GetQueueUrlInput, optFns ...func(*sqs.Options)) (*sqs.GetQueueUrlOutput, error)
}

type AwsSqs struct {
	queueName string
	client    AwsSqsClient
}

func (r *AwsSqs) SetQueueName(queueName string) {
	r.queueName = queueName
}

func (r *AwsSqs) getQueueURL() (*sqs.GetQueueUrlOutput, error) {
	ctx, cancel := context.WithTimeout(context.Background(), connectionTimeout)
	defer cancel()

	queueNameInput := &sqs.GetQueueUrlInput{
		QueueName: &r.queueName,
	}

	result, err := r.client.GetQueueUrl(ctx, queueNameInput)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *AwsSqs) PublishMessage(message string) error {
	ctx, cancel := context.WithTimeout(context.Background(), connectionTimeout)
	defer cancel()

	queue, err := r.getQueueURL()
	if err != nil {
		return err
	}
	params := &sqs.SendMessageInput{
		MessageBody: aws.String(message),
		QueueUrl:    queue.QueueUrl,
	}
	_, err = r.client.SendMessage(ctx, params)
	if err != nil {
		return err
	}

	return nil
}

func NewAwsSqs(client AwsSqsClient) *AwsSqs {
	return &AwsSqs{
		client: client,
	}
}
