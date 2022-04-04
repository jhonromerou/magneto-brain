package internal

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/jhonromerou/magneto-brain/src/domain"
	"github.com/jhonromerou/magneto-brain/src/domain/mocks"
	"github.com/jhonromerou/magneto-brain/src/domain/services"

	"github.com/aws/aws-lambda-go/events"
)

func Test_Handler_Handle(t *testing.T) {
	type fields struct {
		logger     *mocks.Logger
		repository *mocks.QueueAnalysisRepository
	}

	commonFields := fields{
		logger:     &mocks.Logger{},
		repository: &mocks.QueueAnalysisRepository{},
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
			name:   "error converting json",
			fields: commonFields,
			args: args{
				ctx: ctx,
				req: events.APIGatewayProxyRequest{
					Body: ``,
				},
			},
			funcMock: func(f fields) {
				f.logger.On("Info", "Processing analysis with data ").Once()
				dnaSequenceJSON := services.DnaSequenceJSON{}

				err := json.Unmarshal([]byte(""), &dnaSequenceJSON)
				f.logger.On("ErrorWithDetail", err, "error with dna pattern sequence").Once()
			},
			wantError: true,
			want: events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				Body: `{"name":"dna_error","format":"error with dna {dna} pattern sequence","values":{"dna":""}}`,
			},
		},
		{
			name:   "dna pattern sequence invalid",
			fields: commonFields,
			args: args{
				ctx: ctx,
				req: events.APIGatewayProxyRequest{
					Body: `{"dna":["ERROR","FORMAT","ADN","EXAM","PLE"]}`,
				},
			},
			funcMock: func(f fields) {
				f.logger.On("Info", `Processing analysis with data {"dna":["ERROR","FORMAT","ADN","EXAM","PLE"]}`).Once()
				f.logger.On("Error", "dna pattern sequence invalid").Once()
			},
			wantError: true,
			want: events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				Body: `{"name":"dna_patter_sequence_invalid","format":"error validating that dna {dna} have allow pattern sequence","values":{"dna":"ERROR,FORMAT,ADN,EXAM,PLE"}}`,
			},
		},
		{
			name:   "error publish message",
			fields: commonFields,
			args: args{
				ctx: ctx,
				req: events.APIGatewayProxyRequest{
					Body: `{"dna":["CGTCAA", "CTGTGC", "TAATTT", "AGAAGT", "CCGCAA", "TCACTA"]}`,
				},
			},
			funcMock: func(f fields) {
				f.logger.On("Info", `Processing analysis with data {"dna":["CGTCAA", "CTGTGC", "TAATTT", "AGAAGT", "CCGCAA", "TCACTA"]}`).Once()
				message := `{"dna_type":"dnaMutant","dna_sequence":"CGTCAA,CTGTGC,TAATTT,AGAAGT,CCGCAA,TCACTA"}`
				someError := errors.New("some error...")
				f.repository.On("PublishMessage", message).Once().Return(someError)
				f.logger.On("ErrorWithDetail", someError, "error publishing message on the queue").Once()
			},
			wantError: false,
			want: events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				Body: `{"name":"dna_error","format":"error publishing message on the queue"}`,
			},
		},
		{
			name:   "dna not is of mutant",
			fields: commonFields,
			args: args{
				ctx: ctx,
				req: events.APIGatewayProxyRequest{
					Body: `{"dna":["CGACAA","CTGTGC","TAGTTT","AGAGGT","CCGCAA","TCACTA"]}`,
				},
			},
			funcMock: func(f fields) {
				f.logger.On("Info", `Processing analysis with data {"dna":["CGACAA","CTGTGC","TAGTTT","AGAGGT","CCGCAA","TCACTA"]}`).Once()
				message := `{"dna_type":"dnaHuman","dna_sequence":"CGACAA,CTGTGC,TAGTTT,AGAGGT,CCGCAA,TCACTA"}`
				f.repository.On("PublishMessage", message).Once().Return(nil)
			},

			want: events.APIGatewayProxyResponse{
				StatusCode: http.StatusForbidden,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				Body: `{"name":"dna_not_is_of_mutant","format":"request with dna {dna} not is of mutant","values":{"dna":"CGACAA,CTGTGC,TAGTTT,AGAGGT,CCGCAA,TCACTA"}}`,
			},
		},
		{
			name:   "dna is of mutant",
			fields: commonFields,
			args: args{
				ctx: ctx,
				req: events.APIGatewayProxyRequest{
					Body: `{"dna":["CGTCAA", "CTGTGC", "TAATTT", "AGAAGT", "CCGCAA", "TCACTA"]}`,
				},
			},
			funcMock: func(f fields) {
				f.logger.On("Info", `Processing analysis with data {"dna":["CGTCAA", "CTGTGC", "TAATTT", "AGAAGT", "CCGCAA", "TCACTA"]}`).Once()
				message := `{"dna_type":"dnaMutant","dna_sequence":"CGTCAA,CTGTGC,TAATTT,AGAAGT,CCGCAA,TCACTA"}`
				f.repository.On("PublishMessage", message).Once().Return(nil)
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: http.StatusOK,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				Body: `{"name":"dna_is_of_mutant","format":"request with dna {dna} is of mutant","values":{"dna":"CGTCAA,CTGTGC,TAATTT,AGAAGT,CCGCAA,TCACTA"}}`,
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
