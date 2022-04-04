package repositories

import (
	"errors"
	"testing"

	"github.com/jhonromerou/magneto-brain/src/domain"
	"github.com/jhonromerou/magneto-brain/src/domain/mocks"
)

func Test_Analysis_Repository_PublishMessage(t *testing.T) {
	type fields struct {
		queueName     string
		message       string
		repository    *mocks.QueueRepository
		envRepository *mocks.EnvironmentRespository
	}
	tests := []struct {
		name      string
		fields    fields
		funcMock  func(fields)
		wantError bool
	}{
		{
			name: "error publish sqs",
			fields: fields{
				queueName:     "queue-name",
				message:       `{"dnaType":"dnaMutant","dnaSequence":"a,bc,s"}`,
				repository:    &mocks.QueueRepository{},
				envRepository: &mocks.EnvironmentRespository{},
			},
			funcMock: func(f fields) {
				f.envRepository.On("Get", "QUEUE_NAME_ANALYSIS").Once().Return(f.queueName)
				f.repository.On("SetQueueName", f.queueName).Once()
				err := errors.New("error publish message")
				f.repository.On("PublishMessage", f.message).Once().Return(err)
			},
			wantError: true,
		},
		{
			name: "success publish message",
			fields: fields{
				queueName:     "queue-name",
				message:       `{"dnaType":"dnaMutant","dnaSequence":"a,bc,s"}`,
				repository:    &mocks.QueueRepository{},
				envRepository: &mocks.EnvironmentRespository{},
			},
			funcMock: func(f fields) {
				f.envRepository.On("Get", "QUEUE_NAME_ANALYSIS").Once().Return(f.queueName)
				f.repository.On("SetQueueName", f.queueName).Once()
				f.repository.On("PublishMessage", f.message).Once().Return(nil)
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.funcMock(tt.fields)
			repository := NewQueueAnalysisRepository(tt.fields.repository, tt.fields.envRepository)
			err := repository.PublishMessage(tt.fields.message)
			if err != nil != tt.wantError {
				t.Fatal(domain.TestingErrorNotMatched(1, err, tt.wantError))
			}
		})
	}
}
