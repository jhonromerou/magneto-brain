package repositories

import (
	"github.com/jhonromerou/magneto-brain/src/domain"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
)

type DynamodbStatsRepository struct {
	tableName    string
	rows         []domain.StatsModel
	dbConnection domain.DatabaseRepository
}

func (r *DynamodbStatsRepository) SetDnaAnalysis(item domain.StatsModel) error {
	queryBuilder := r.dbConnection.Query()
	queryBuilder.SetTable(r.tableName)
	item.Group = domain.STATS_GROUP_DNA_ANALIZE
	err := r.dbConnection.Upsert(item)
	if err != nil {
		return err
	}

	return nil
}

func (r *DynamodbStatsRepository) GetDnaAnalysisByName(name string) (domain.StatsModel, error) {
	emptyResult := domain.StatsModel{}
	finders := domain.QueryFinders{
		[3]string{"group", "=", domain.STATS_GROUP_DNA_ANALIZE},
		[3]string{"name", "=", name},
	}
	query := r.dbConnection.Query()
	query.SetTable(r.tableName)
	query.Fields("*")
	query.Where(finders)
	output, err := query.Get()
	if err != nil {
		return emptyResult, err
	}

	rows := r.rows
	err = attributevalue.UnmarshalListOfMaps(output.Items, &rows)
	if err != nil {
		return emptyResult, err
	}

	if len(rows) == 0 {
		return emptyResult, nil
	}

	return rows[0], nil
}

func (r *DynamodbStatsRepository) NewDnaAnalysisWithName(name string) error {
	actualStat, err := r.GetDnaAnalysisByName(name)
	if err != nil {
		return err
	}
	newStat := actualStat
	if newStat.Name == "" {
		newStat.Name = name
	}
	newStat.Quantity = actualStat.Quantity + 1

	err = r.SetDnaAnalysis(newStat)
	if err != nil {
		return err
	}

	return nil
}

func NewDynamodbStatsRepository(c domain.DatabaseRepository, envs domain.EnvironmentRespository) *DynamodbStatsRepository {
	return &DynamodbStatsRepository{
		tableName:    envs.Get("DB_TABLE_STATS"),
		dbConnection: c,
	}
}
