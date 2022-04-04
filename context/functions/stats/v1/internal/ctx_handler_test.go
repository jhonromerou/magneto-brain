package internal

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/jhonromerou/magneto-brain/src/domain"
	"github.com/jhonromerou/magneto-brain/src/domain/mocks"

	"github.com/aws/aws-lambda-go/events"
)

func Test_Handler_Handle(t *testing.T) {
	type fields struct {
		logger     *mocks.Logger
		repository *mocks.StatsRepository
	}

	commonFields := fields{
		logger:     &mocks.Logger{},
		repository: &mocks.StatsRepository{},
	}

	type args struct {
		ctx context.Context
		req events.APIGatewayProxyRequest
	}

	ctx := context.Background()

	tests := []struct {
		name      string
		fields    fields
		args      args
		funcMock  func(fields)
		wantError bool
		want      events.APIGatewayProxyResponse
	}{
		{
			name:   "error getting analysis of type human",
			fields: commonFields,
			args: args{
				ctx: ctx,
				req: events.APIGatewayProxyRequest{},
			},
			funcMock: func(f fields) {
				f.logger.On("Info", "Getting stats").Once()
				expectedMessage := "error getting dnaHuman stats"
				err := errors.New(expectedMessage)
				f.repository.On("GetDnaAnalysisByName", domain.STATS_NAME_DNA_HUMAN).Once().Return(domain.StatsModel{}, err)
				f.logger.On("ErrorWithDetail", err, expectedMessage).Once()
			},
			wantError: true,
			want: events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				Body: `{}`,
			},
		},
		{
			name:   "error getting analysis of type mutant",
			fields: commonFields,
			args: args{
				ctx: ctx,
				req: events.APIGatewayProxyRequest{},
			},
			funcMock: func(f fields) {
				f.logger.On("Info", "Getting stats").Once()
				expectedMessage := "error getting dnaMutant stats"
				err := errors.New(expectedMessage)
				f.repository.On("GetDnaAnalysisByName", domain.STATS_NAME_DNA_HUMAN).Once().Return(domain.StatsModel{}, nil)
				f.repository.On("GetDnaAnalysisByName", domain.STATS_NAME_DNA_MUTANT).Once().Return(domain.StatsModel{}, err)
				f.logger.On("ErrorWithDetail", err, expectedMessage).Once()
			},
			wantError: true,
			want: events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				Body: `{}`,
			},
		},
		{
			name:   "success report",
			fields: commonFields,
			args: args{
				ctx: ctx,
				req: events.APIGatewayProxyRequest{},
			},
			funcMock: func(f fields) {
				f.logger.On("Info", "Getting stats").Once()
				f.repository.On("GetDnaAnalysisByName", domain.STATS_NAME_DNA_HUMAN).Once().Return(domain.StatsModel{Quantity: 1}, nil)
				f.repository.On("GetDnaAnalysisByName", domain.STATS_NAME_DNA_MUTANT).Once().Return(domain.StatsModel{Quantity: 2}, nil)
			},
			wantError: false,
			want: events.APIGatewayProxyResponse{
				StatusCode: http.StatusOK,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				Body: `{"dna":{"count_mutant_dna":2,"count_human_dna":1,"ratio":0.66}}`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.funcMock(tt.fields)
			handler := NewHandler(tt.fields.logger, tt.fields.repository)

			got, err := handler.Handle(tt.args.ctx, tt.args.req)
			if (err != nil) && tt.wantError {
				t.Fatal(domain.TestingErrorNotMatched(1, got, tt.want))
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf(domain.TestingErrorNotMatched(2, got, tt.want))
			}
		})
	}
}
