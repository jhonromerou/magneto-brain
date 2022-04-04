package repositories

import (
	"github.com/jhonromerou/magneto-brain/src/domain"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
)

type DynamodbAnalysisRepository struct {
	tableName    string
	rows         []domain.AnalysisModel
	dbConnection domain.DatabaseRepository
}

func (r *DynamodbAnalysisRepository) Register(item domain.AnalysisModel) error {
	r.dbConnection.SetTable(r.tableName)
	err := r.dbConnection.Upsert(item)

	return err
}

func (r *DynamodbAnalysisRepository) GetAnalysisResult(dnaType string, dnaSequence string) (domain.AnalysisModel, error) {
	emptyResult := domain.AnalysisModel{}

	finders := domain.QueryFinders{
		[3]string{"dna_type", "=", dnaType},
		[3]string{"dna_sequence", "=", dnaSequence},
	}

	r.dbConnection.SetTable(r.tableName)
	r.dbConnection.Fields("*")
	r.dbConnection.Where(finders)
	output, err := r.dbConnection.Get()
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

func NewDynamodbAnalysisRepository(c domain.DatabaseRepository, envs domain.EnvironmentRespository) *DynamodbAnalysisRepository {
	return &DynamodbAnalysisRepository{
		tableName:    envs.Get("DB_TABLE_ANALYSIS"),
		dbConnection: c,
	}
}
