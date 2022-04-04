package repositories

import (
	"github.com/jhonromerou/magneto-brain/src/domain"
)

type AwsSqsAnalysisRepository struct {
	queueName  string
	repository domain.QueueRepository
}

func (q *AwsSqsAnalysisRepository) PublishMessage(message string) error {
	q.repository.SetQueueName(q.queueName)
	return q.repository.PublishMessage(message)
}

func NewQueueAnalysisRepository(r domain.QueueRepository, envs domain.EnvironmentRespository) *AwsSqsAnalysisRepository {
	return &AwsSqsAnalysisRepository{
		queueName:  envs.Get("QUEUE_NAME_ANALYSIS"),
		repository: r,
	}
}
