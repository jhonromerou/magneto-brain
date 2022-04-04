package awsdynamodb

import (
	"errors"
	"testing"

	"github.com/jhonromerou/magneto-brain/src/domain"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/stretchr/testify/mock"
)

func Test_Aws_Dynamodb_Get(t *testing.T) {
	type fields struct {
		tableName   string
		awsDynamodb *MockAwsDynamoDbClient
	}
	commomQuery := `SELECT * FROM "table-name" WHERE "field1" = 'value1' AND "field2" = 'value2'`
	tests := []struct {
		name      string
		fields    fields
		funcMock  func(fields)
		wantError bool
		want      interface{}
	}{
		{
			name: "error get rows",
			fields: fields{
				tableName:   "table-name",
				awsDynamodb: &MockAwsDynamoDbClient{},
			},
			funcMock: func(f fields) {
				stament := &dynamodb.ExecuteStatementInput{
					Statement: aws.String(commomQuery),
				}
				err := errors.New("some error")
				f.awsDynamodb.On("ExecuteStatement", mock.Anything, stament).Return(&dynamodb.ExecuteStatementOutput{}, err)
			},
			wantError: true,
		},
		{
			name: "success get rows",
			fields: fields{
				tableName:   "table-name",
				awsDynamodb: &MockAwsDynamoDbClient{},
			},
			funcMock: func(f fields) {
				stament := &dynamodb.ExecuteStatementInput{
					Statement: aws.String(commomQuery),
				}
				f.awsDynamodb.On("ExecuteStatement", mock.Anything, stament).Return(&dynamodb.ExecuteStatementOutput{}, nil)
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.funcMock(tt.fields)
			awsDynamodb := NewAwsDynamoDb(tt.fields.awsDynamodb)
			awsDynamodb.Query()
			awsDynamodb.SetTable("table-name")
			awsDynamodb.Fields("*")
			awsDynamodb.Where(domain.QueryFinders{
				[3]string{"field1", "=", "value1"},
				[3]string{"field2", "=", "value2"},
			})
			_, err := awsDynamodb.Get()
			if (err != nil) != tt.wantError {
				t.Fatal(domain.TestingErrorNotMatched(1, err, tt.wantError))
			}
		})
	}
}

func Test_Aws_Dynamodb_Upsert(t *testing.T) {
	type fields struct {
		tableName   string
		awsDynamodb *MockAwsDynamoDbClient
	}
	tests := []struct {
		name      string
		fields    fields
		funcMock  func(fields)
		wantError bool
		want      interface{}
	}{
		{
			name: "error upsert information",
			fields: fields{
				tableName:   "table-name",
				awsDynamodb: &MockAwsDynamoDbClient{},
			},
			funcMock: func(f fields) {
				itemInformation := domain.AnalysisModel{
					DnaType:     "human",
					DnaSequence: "abc,deg,bdd",
				}
				attributes, _ := attributevalue.MarshalMap(itemInformation)
				stament := &dynamodb.PutItemInput{
					TableName: aws.String("table-name"),
					Item:      attributes,
				}
				err := errors.New("some error")
				f.awsDynamodb.On("PutItem", mock.Anything, stament).Return(&dynamodb.PutItemOutput{}, err)
			},
			wantError: true,
		},
		{
			name: "success saving information",
			fields: fields{
				tableName:   "table-name",
				awsDynamodb: &MockAwsDynamoDbClient{},
			},
			funcMock: func(f fields) {
				itemInformation := domain.AnalysisModel{
					DnaType:     "human",
					DnaSequence: "abc,deg,bdd",
				}
				attributes, _ := attributevalue.MarshalMap(itemInformation)
				stament := &dynamodb.PutItemInput{
					TableName: aws.String("table-name"),
					Item:      attributes,
				}
				f.awsDynamodb.On("PutItem", mock.Anything, stament).Return(&dynamodb.PutItemOutput{}, nil)
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.funcMock(tt.fields)
			awsDynamodb := NewAwsDynamoDb(tt.fields.awsDynamodb)
			itemInformation := domain.AnalysisModel{
				DnaType:     "human",
				DnaSequence: "abc,deg,bdd",
			}
			awsDynamodb.SetTable(tt.fields.tableName)
			err := awsDynamodb.Upsert(itemInformation)
			if (err != nil) != tt.wantError {
				t.Fatal(domain.TestingErrorNotMatched(1, err, tt.wantError))
			}

			tt.fields.awsDynamodb.AssertExpectations(t)
		})
	}
}
