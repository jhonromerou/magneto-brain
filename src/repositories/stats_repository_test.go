package repositories

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/jhonromerou/magneto-brain/src/domain"
	"github.com/jhonromerou/magneto-brain/src/domain/mocks"
)

func Test_Analysis_Repository_SetDnaAnalysis(t *testing.T) {
	type fields struct {
		tableName     string
		item          domain.StatsModel
		dbConnection  *mocks.DatabaseRepository
		envRepository *mocks.EnvironmentRespository
	}
	tests := []struct {
		name      string
		fields    fields
		funcMock  func(fields)
		wantError bool
	}{
		{
			name: "error set stats",
			fields: fields{
				tableName: "table",
				item: domain.StatsModel{
					Group:    domain.STATS_GROUP_DNA_ANALIZE,
					Name:     "dnaMutant",
					Quantity: 1,
				},
				dbConnection:  &mocks.DatabaseRepository{},
				envRepository: &mocks.EnvironmentRespository{},
			},
			funcMock: func(f fields) {
				f.envRepository.On("Get", "DB_TABLE_STATS").Once().Return("table")
				f.dbConnection.On("Query").Once().Return(f.dbConnection)
				f.dbConnection.On("SetTable", "table").Once()
				err := errors.New("error in repository")
				f.dbConnection.On("Upsert", f.item).Once().Return(err)
			},
			wantError: true,
		},
		{
			name: "success set stats",
			fields: fields{
				tableName: "table",
				item: domain.StatsModel{
					Group:    domain.STATS_GROUP_DNA_ANALIZE,
					Name:     "dnaMutant",
					Quantity: 1,
				},
				dbConnection:  &mocks.DatabaseRepository{},
				envRepository: &mocks.EnvironmentRespository{},
			},
			funcMock: func(f fields) {
				f.envRepository.On("Get", "DB_TABLE_STATS").Once().Return("table")
				f.dbConnection.On("Query").Once().Return(f.dbConnection)
				f.dbConnection.On("SetTable", "table").Once()
				f.dbConnection.On("Upsert", f.item).Once().Return(nil)
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.funcMock(tt.fields)
			repository := NewDynamodbStatsRepository(tt.fields.dbConnection, tt.fields.envRepository)
			err := repository.SetDnaAnalysis(tt.fields.item)
			if err != nil != tt.wantError {
				t.Fatal(domain.TestingErrorNotMatched(1, err, tt.wantError))
			}
		})
	}
}

func Test_Analysis_Repository_GetDnaAnalysisByName(t *testing.T) {
	type fields struct {
		tableName     string
		dnaType       string
		dbConnection  *mocks.DatabaseRepository
		envRepository *mocks.EnvironmentRespository
	}
	tests := []struct {
		name      string
		fields    fields
		funcMock  func(fields)
		wantError bool
	}{
		{
			name: "error get dna analysis stats by name",
			fields: fields{
				tableName:     "table",
				dnaType:       "dnaMutant",
				dbConnection:  &mocks.DatabaseRepository{},
				envRepository: &mocks.EnvironmentRespository{},
			},
			funcMock: func(f fields) {
				f.envRepository.On("Get", "DB_TABLE_STATS").Once().Return("table")
				f.dbConnection.On("Query").Once().Return(f.dbConnection)
				f.dbConnection.On("SetTable", "table").Once()
				f.dbConnection.On("Fields", "*").Once()
				finders := domain.QueryFinders{
					[3]string{"group", "=", domain.STATS_GROUP_DNA_ANALIZE},
					[3]string{"name", "=", f.dnaType},
				}
				f.dbConnection.On("Where", finders).Once()
				err := errors.New("error in repository")
				f.dbConnection.On("Get").Once().Return(&dynamodb.ExecuteStatementOutput{}, err)
			},
			wantError: true,
		},
		{
			name: "success get analysis",
			fields: fields{
				tableName:     "table",
				dnaType:       "dnaMutant",
				dbConnection:  &mocks.DatabaseRepository{},
				envRepository: &mocks.EnvironmentRespository{},
			},
			funcMock: func(f fields) {
				f.envRepository.On("Get", "DB_TABLE_STATS").Once().Return("table")
				f.dbConnection.On("Query").Once().Return(f.dbConnection)
				f.dbConnection.On("SetTable", "table").Once()
				f.dbConnection.On("Fields", "*").Once()
				finders := domain.QueryFinders{
					[3]string{"group", "=", domain.STATS_GROUP_DNA_ANALIZE},
					[3]string{"name", "=", f.dnaType},
				}
				f.dbConnection.On("Where", finders).Once()
				f.dbConnection.On("Get").Once().Return(&dynamodb.ExecuteStatementOutput{}, nil)
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.funcMock(tt.fields)
			repository := NewDynamodbStatsRepository(tt.fields.dbConnection, tt.fields.envRepository)
			_, err := repository.GetDnaAnalysisByName(tt.fields.dnaType)
			if err != nil != tt.wantError {
				t.Error(domain.TestingErrorNotMatched(1, err, tt.wantError))
			}
		})
	}
}

func Test_Analysis_Repository_NewDnaAnalysisWithName(t *testing.T) {
	type fields struct {
		tableName     string
		dnaType       string
		dbConnection  *mocks.DatabaseRepository
		envRepository *mocks.EnvironmentRespository
	}
	tests := []struct {
		name      string
		fields    fields
		funcMock  func(fields)
		wantError bool
	}{
		{
			name: "error get dna analysis stats by name",
			fields: fields{
				tableName:     "table",
				dnaType:       "dnaMutant",
				dbConnection:  &mocks.DatabaseRepository{},
				envRepository: &mocks.EnvironmentRespository{},
			},
			funcMock: func(f fields) {
				f.envRepository.On("Get", "DB_TABLE_STATS").Once().Return("table")
				f.dbConnection.On("Query").Once().Return(f.dbConnection)
				f.dbConnection.On("SetTable", "table").Once()
				f.dbConnection.On("Fields", "*").Once()
				finders := domain.QueryFinders{
					[3]string{"group", "=", domain.STATS_GROUP_DNA_ANALIZE},
					[3]string{"name", "=", f.dnaType},
				}
				f.dbConnection.On("Where", finders).Once()
				err := errors.New("error in repository")
				f.dbConnection.On("Get").Once().Return(&dynamodb.ExecuteStatementOutput{}, err)
			},
			wantError: true,
		},
		{
			name: "error set dna analysis stats by name",
			fields: fields{
				tableName:     "table",
				dnaType:       "dnaMutant",
				dbConnection:  &mocks.DatabaseRepository{},
				envRepository: &mocks.EnvironmentRespository{},
			},
			funcMock: func(f fields) {
				f.envRepository.On("Get", "DB_TABLE_STATS").Once().Return("table")
				f.dbConnection.On("Query").Once().Return(f.dbConnection)
				f.dbConnection.On("SetTable", "table").Once()
				f.dbConnection.On("Fields", "*").Once()
				finders := domain.QueryFinders{
					[3]string{"group", "=", domain.STATS_GROUP_DNA_ANALIZE},
					[3]string{"name", "=", f.dnaType},
				}
				f.dbConnection.On("Where", finders).Once()
				f.dbConnection.On("Get").Once().Return(&dynamodb.ExecuteStatementOutput{}, nil)
				f.dbConnection.On("Query").Once().Return(f.dbConnection)
				f.dbConnection.On("SetTable", "table").Once()
				err := errors.New("error saving stats")
				itemSave := domain.StatsModel{
					Group:    domain.STATS_GROUP_DNA_ANALIZE,
					Name:     domain.STATS_NAME_DNA_MUTANT,
					Quantity: 1,
				}
				f.dbConnection.On("Upsert", itemSave).Once().Return(err)
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.funcMock(tt.fields)
			repository := NewDynamodbStatsRepository(tt.fields.dbConnection, tt.fields.envRepository)
			err := repository.NewDnaAnalysisWithName(tt.fields.dnaType)
			if err != nil != tt.wantError {
				t.Error(domain.TestingErrorNotMatched(1, err, tt.wantError))
			}
		})
	}
}
