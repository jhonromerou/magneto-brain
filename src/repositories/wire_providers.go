package repositories

import (
	"github.com/jhonromerou/magneto-brain/src/domain"
	"github.com/jhonromerou/magneto-brain/src/infrastructure/awsdynamodb"
	"github.com/jhonromerou/magneto-brain/src/infrastructure/awssqs"
	"github.com/jhonromerou/magneto-brain/src/infrastructure/dotenv"

	"github.com/google/wire"
)

// SetStatsRepository defines dependency injection to create stats repository instance.
var SetStatsRepository = wire.NewSet(
	dotenv.SetEnviroments,
	awsdynamodb.SetAwsDynamodb,
	NewDynamodbStatsRepository,
	wire.Bind(new(domain.StatsRepository), new(*DynamodbStatsRepository)),
)

var SetAnalysisAndStatsRepository = wire.NewSet(
	dotenv.SetEnviroments,
	awsdynamodb.SetAwsDynamodb,
	NewDynamodbAnalysisRepository,
	NewDynamodbStatsRepository,
	wire.Bind(new(domain.AnalysisRepository), new(*DynamodbAnalysisRepository)),
	wire.Bind(new(domain.StatsRepository), new(*DynamodbStatsRepository)),
)

// SetAnalysisRepository defines dependency injection to create analysis repository instance.
var SetAnalysisRepository = wire.NewSet(
	dotenv.SetEnviroments,
	awsdynamodb.SetAwsDynamodb,
	NewDynamodbAnalysisRepository,
	wire.Bind(new(domain.AnalysisRepository), new(*DynamodbAnalysisRepository)),
)

// SetQueueAnalysisRepository defines dependency injection to create queue analysis repository instance.
var SetQueueAnalysisRepository = wire.NewSet(
	dotenv.SetEnviroments,
	awssqs.SetAwsSqs,
	NewQueueAnalysisRepository,
	wire.Bind(new(domain.QueueRepository), new(*awssqs.AwsSqs)),
	wire.Bind(new(domain.QueueAnalysisRepository), new(*AwsSqsAnalysisRepository)),
)
