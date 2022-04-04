package domain

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// Represents finders to make query like [3]{"fieldName", "condition", "value"}
type QueryFinders [][3]string

type DatabaseInsertOne interface{}

// QueueRepository abstracts common conection to queue.
type QueueRepository interface {
	SetQueueName(queueName string)
	PublishMessage(message string) error
}

type QueueAnalysisRepository interface {
	PublishMessage(message string) error
}

type StatsRepository interface {
	SetDnaAnalysis(item StatsModel) error
	GetDnaAnalysisByName(name string) (StatsModel, error)
	NewDnaAnalysisWithName(name string) error
}

type AnalysisRepository interface {
	Register(ites AnalysisModel) error
	GetAnalysisResult(dnaType string, dnaSequence string) (AnalysisModel, error)
}

type DatabaseRepository interface {
	Query() DatabaseRepository
	SetTable(tableName string)
	Fields(fields string)
	Where(finders QueryFinders)
	// TODO: desacoplar de dynamodb
	Get() (*dynamodb.ExecuteStatementOutput, error)
	Upsert(DatabaseInsertOne) error
}
