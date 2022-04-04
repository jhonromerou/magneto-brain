package awsdynamodb

import (
	"context"
	"fmt"
	"time"

	"github.com/jhonromerou/magneto-brain/src/domain"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type AwsDynamoDbClient interface {
	ExecuteStatement(ctx context.Context, params *dynamodb.ExecuteStatementInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ExecuteStatementOutput, error)
	PutItem(context.Context, *dynamodb.PutItemInput, ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
}

type AwsDynamoDb struct {
	tableName     string
	queryString   string
	querySelected bool
	client        AwsDynamoDbClient
}

const connectionTimeout = 25 * time.Second

// SetTable sets table name to connect in dynamodb.
func (r *AwsDynamoDb) SetTable(tableName string) {
	r.tableName = tableName
}

func (r *AwsDynamoDb) Query() domain.DatabaseRepository {
	r.tableName = ""
	r.querySelected = false
	r.queryString = ""

	return r
}

// Fields sets specific fields to get in the query.
func (r *AwsDynamoDb) Fields(fields string) {
	if !r.querySelected {
		r.querySelected = true
		r.queryString = fmt.Sprintf(`SELECT %s FROM "%s"`, fields, r.tableName)
	}
}

// Where sets finders/criteria to make a query.
func (r *AwsDynamoDb) Where(finders domain.QueryFinders) {
	var finderBuilder string

	if len(finders) > 0 {
		firstFinder := finders[0]
		finders = finders[1:]
		finderBuilder = fmt.Sprintf(`WHERE "%s" %s '%s'`, firstFinder[0], firstFinder[1], firstFinder[2])
	}
	for _, finder := range finders {
		finderBuilder = fmt.Sprintf(`%s AND "%s" %s '%s'`, finderBuilder, finder[0], finder[1], finder[2])
	}

	r.queryString = fmt.Sprintf(`%s %s`, r.queryString, finderBuilder)
}

// Get gets rows of query.
func (r *AwsDynamoDb) Get() (*dynamodb.ExecuteStatementOutput, error) {
	ctx, cancel := context.WithTimeout(context.Background(), connectionTimeout)
	defer cancel()
	output, err := r.client.ExecuteStatement(ctx, &dynamodb.ExecuteStatementInput{
		Statement: aws.String(r.queryString),
	})

	if err != nil {
		return nil, err
	}

	return output, nil
}

// Upsert inserts one item on dispatch queue
func (r *AwsDynamoDb) Upsert(item domain.DatabaseInsertOne) error {
	ctx, cancel := context.WithTimeout(context.Background(), connectionTimeout)
	defer cancel()

	attributes, err := attributevalue.MarshalMap(item)
	if err != nil {
		return err
	}

	_, err = r.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &r.tableName,
		Item:      attributes,
	})

	return err
}

// NewAwsDynamoDb Constructor to AwsDynamoDb.
func NewAwsDynamoDb(client AwsDynamoDbClient) *AwsDynamoDb {
	return &AwsDynamoDb{
		client: client,
	}
}
