package repositories

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/jhonromerou/magneto-brain/src/domain"
	"github.com/jhonromerou/magneto-brain/src/domain/mocks"
)

func Test_Analysis_Repository_Register(t *testing.T) {
	type fields struct {
		tableName     string
		item          domain.AnalysisModel
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
			name: "error when register",
			fields: fields{
				tableName: "table",
				item: domain.AnalysisModel{
					DnaType:     "dnaMutant",
					DnaSequence: "savb,,s,ws,,",
				},
				dbConnection:  &mocks.DatabaseRepository{},
				envRepository: &mocks.EnvironmentRespository{},
			},
			funcMock: func(f fields) {
				f.envRepository.On("Get", "DB_TABLE_ANALYSIS").Once().Return("table")
				f.dbConnection.On("SetTable", "table").Once()
				err := errors.New("error in repository")
				f.dbConnection.On("Upsert", f.item).Once().Return(err)
			},
			wantError: true,
		},
		{
			name: "success register",
			fields: fields{
				tableName: "table",
				item: domain.AnalysisModel{
					DnaType:     "dnaMutant",
					DnaSequence: "savb,,s,ws,,",
				},
				dbConnection:  &mocks.DatabaseRepository{},
				envRepository: &mocks.EnvironmentRespository{},
			},
			funcMock: func(f fields) {
				f.envRepository.On("Get", "DB_TABLE_ANALYSIS").Once().Return("table")
				f.dbConnection.On("SetTable", "table").Once()
				f.dbConnection.On("Upsert", f.item).Once().Return(nil)
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.funcMock(tt.fields)
			repository := NewDynamodbAnalysisRepository(tt.fields.dbConnection, tt.fields.envRepository)
			err := repository.Register(tt.fields.item)
			if err != nil != tt.wantError {
				t.Fatal(domain.TestingErrorNotMatched(1, err, tt.wantError))
			}
		})
	}
}

func Test_Analysis_Repository_GetAnalysisResult(t *testing.T) {
	type fields struct {
		tableName     string
		dnaType       string
		dnaSequence   string
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
			name: "error get analysis result",
			fields: fields{
				tableName:     "table",
				dnaType:       "dnaMutant",
				dnaSequence:   "abc,dfc",
				dbConnection:  &mocks.DatabaseRepository{},
				envRepository: &mocks.EnvironmentRespository{},
			},
			funcMock: func(f fields) {
				f.envRepository.On("Get", "DB_TABLE_ANALYSIS").Once().Return("table")
				f.dbConnection.On("Query").Once().Return(f.dbConnection)
				f.dbConnection.On("SetTable", "table").Once()
				f.dbConnection.On("Fields", "*").Once()
				finders := domain.QueryFinders{
					[3]string{"dna_type", "=", f.dnaType},
					[3]string{"dna_sequence", "=", f.dnaSequence},
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
				dnaSequence:   "abc,dfc",
				dbConnection:  &mocks.DatabaseRepository{},
				envRepository: &mocks.EnvironmentRespository{},
			},
			funcMock: func(f fields) {
				f.envRepository.On("Get", "DB_TABLE_ANALYSIS").Once().Return("table")
				f.dbConnection.On("Query").Once().Return(f.dbConnection)
				f.dbConnection.On("SetTable", "table").Once()
				f.dbConnection.On("Fields", "*").Once()
				finders := domain.QueryFinders{
					[3]string{"dna_type", "=", f.dnaType},
					[3]string{"dna_sequence", "=", f.dnaSequence},
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
			repository := NewDynamodbAnalysisRepository(tt.fields.dbConnection, tt.fields.envRepository)
			_, err := repository.GetAnalysisResult(tt.fields.dnaType, tt.fields.dnaSequence)
			if err != nil != tt.wantError {
				t.Error(domain.TestingErrorNotMatched(1, err, tt.wantError))
			}
		})
	}
}
