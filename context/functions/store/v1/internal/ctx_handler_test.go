package internal

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/jhonromerou/magneto-brain/src/domain"
	"github.com/jhonromerou/magneto-brain/src/domain/mocks"

	"github.com/aws/aws-lambda-go/events"
)

func Test_Handler_Handle(t *testing.T) {
	type fields struct {
		logger             *mocks.Logger
		analysisRepository *mocks.AnalysisRepository
		statsRepository    *mocks.StatsRepository
	}

	commonFields := fields{
		logger:             &mocks.Logger{},
		analysisRepository: &mocks.AnalysisRepository{},
		statsRepository:    &mocks.StatsRepository{},
	}

	commonAnalysis := domain.AnalysisModel{
		DnaType:     "dnaHuman",
		DnaSequence: "xxxx",
	}

	type args struct {
		ctx context.Context
		req events.SQSEvent
	}

	ctx := context.Background()

	tests := []struct {
		name      string
		fields    fields
		args      args
		funcMock  func(fields)
		wantError bool
		want      error
	}{
		{
			name:   "without message",
			fields: commonFields,
			args: args{
				ctx: ctx,
				req: events.SQSEvent{
					Records: []events.SQSMessage{},
				},
			},
			funcMock: func(f fields) {
				f.logger.On("Info", "Processing 0 queue messages").Once()
			},
			wantError: true,
			want:      nil,
		},
		{
			name:   "request dna type value on message",
			fields: commonFields,
			args: args{
				ctx: ctx,
				req: events.SQSEvent{
					Records: []events.SQSMessage{
						{
							Body: `{"dna_type":""}`,
						},
					},
				},
			},
			funcMock: func(f fields) {
				f.logger.On("Info", "Processing 1 queue messages").Once()
				f.logger.On("Info", `#0 Message: {"dna_type":""}`)
				err := errors.New("field dna_type is invalid or empty")
				f.logger.On("ErrorWithDetail", err, "cannot parse message")
			},
			wantError: true,
			want:      nil,
		},
		{
			name:   "request dna sequence value on message",
			fields: commonFields,
			args: args{
				ctx: ctx,
				req: events.SQSEvent{
					Records: []events.SQSMessage{
						{
							Body: `{"dna_type":"dnaMutant","dna_sequence":""}`,
						},
					},
				},
			},
			funcMock: func(f fields) {
				f.logger.On("Info", "Processing 1 queue messages").Once()
				f.logger.On("Info", `#0 Message: {"dna_type":"dnaMutant","dna_sequence":""}`)
				err := errors.New("field dna_sequence is invalid or empty")
				f.logger.On("ErrorWithDetail", err, "cannot parse message")
			},
			wantError: true,
			want:      nil,
		},
		{
			name:   "error with parse message",
			fields: commonFields,
			args: args{
				ctx: ctx,
				req: events.SQSEvent{
					Records: []events.SQSMessage{
						{
							Body: `@@`,
						},
					},
				},
			},
			funcMock: func(f fields) {
				f.logger.On("Info", "Processing 1 queue messages").Once()
				f.logger.On("Info", "#0 Message: @@")
				message := ResultMessage{}
				err := json.Unmarshal([]byte("@@"), &message)
				f.logger.On("ErrorWithDetail", err, "cannot parse message").Once()
			},
			wantError: true,
			want:      nil,
		},
		{
			name:   "error get anlysis result",
			fields: commonFields,
			args: args{
				ctx: ctx,
				req: events.SQSEvent{
					Records: []events.SQSMessage{
						{
							Body: `{"dna_type":"dnaHuman","dna_sequence":"xxxx"}`,
						},
					},
				},
			},
			funcMock: func(f fields) {
				f.logger.On("Info", "Processing 1 queue messages").Once()
				f.logger.On("Info", `#0 Message: {"dna_type":"dnaHuman","dna_sequence":"xxxx"}`)
				result := ResultMessage{}
				json.Unmarshal([]byte(`{"dna_type":"dnaHuman","dna_sequence":"xxxx"}`), &result)
				err := errors.New("some error getting analysis")
				f.logger.On("ErrorWithDetail", err, "error validate that analysis exists").Once()
				f.analysisRepository.On("GetAnalysisResult", result.DnaType, result.DnaSequence).Once().Return(domain.AnalysisModel{}, err)
			},
			wantError: true,
			want:      nil,
		},
		{
			name:   "analysis already exists",
			fields: commonFields,
			args: args{
				ctx: ctx,
				req: events.SQSEvent{
					Records: []events.SQSMessage{
						{
							Body: `{"dna_type":"dnaHuman","dna_sequence":"xxxx"}`,
						},
					},
				},
			},
			funcMock: func(f fields) {
				f.logger.On("Info", "Processing 1 queue messages").Once()
				f.logger.On("Info", `#0 Message: {"dna_type":"dnaHuman","dna_sequence":"xxxx"}`)
				result := ResultMessage{}
				json.Unmarshal([]byte(`{"dna_type":"dnaHuman","dna_sequence":"xxxx"}`), &result)
				f.analysisRepository.On("GetAnalysisResult", result.DnaType, result.DnaSequence).Once().Return(commonAnalysis, nil)
				f.logger.On("Info", "analysis already exist")
			},
			wantError: true,
			want:      nil,
		},
		{
			name:   "error saving analysis",
			fields: commonFields,
			args: args{
				ctx: ctx,
				req: events.SQSEvent{
					Records: []events.SQSMessage{
						{
							Body: `{"dna_type":"dnaHuman","dna_sequence":"xxxx"}`,
						},
					},
				},
			},
			funcMock: func(f fields) {
				f.logger.On("Info", "Processing 1 queue messages").Once()
				f.logger.On("Info", `#0 Message: {"dna_type":"dnaHuman","dna_sequence":"xxxx"}`)
				result := ResultMessage{}
				json.Unmarshal([]byte(`{"dna_type":"dnaHuman","dna_sequence":"xxxx"}`), &result)
				f.analysisRepository.On("GetAnalysisResult", result.DnaType, result.DnaSequence).Once().Return(domain.AnalysisModel{}, nil)
				err := errors.New("some error saving analysis")
				f.analysisRepository.On("Register", commonAnalysis).Once().Return(err)
				f.logger.On("ErrorWithDetail", err, "error saving analysis result").Once()
			},
			wantError: true,
			want:      nil,
		},
		{
			name:   "error saving stats",
			fields: commonFields,
			args: args{
				ctx: ctx,
				req: events.SQSEvent{
					Records: []events.SQSMessage{
						{
							Body: `{"dna_type":"dnaHuman","dna_sequence":"xxxx"}`,
						},
					},
				},
			},
			funcMock: func(f fields) {
				f.logger.On("Info", "Processing 1 queue messages").Once()
				f.logger.On("Info", `#0 Message: {"dna_type":"dnaHuman","dna_sequence":"xxxx"}`)
				result := ResultMessage{}
				json.Unmarshal([]byte(`{"dna_type":"dnaHuman","dna_sequence":"xxxx"}`), &result)
				f.analysisRepository.On("GetAnalysisResult", result.DnaType, result.DnaSequence).Once().Return(domain.AnalysisModel{}, nil)
				f.analysisRepository.On("Register", commonAnalysis).Once().Return(nil)
				f.logger.On("Info", "Analysis result saved successfully").Once()
				err := errors.New("some error saving stats")
				f.statsRepository.On("NewDnaAnalysisWithName", result.DnaType).Once().Return(err)
				f.logger.On("ErrorWithDetail", err, "error updating stats of analysis dnaHuman").Once()
			},
			wantError: true,
			want:      nil,
		},
		{
			name:   "success saving analysis and stats",
			fields: commonFields,
			args: args{
				ctx: ctx,
				req: events.SQSEvent{
					Records: []events.SQSMessage{
						{
							Body: `{"dna_type":"dnaHuman","dna_sequence":"xxxx"}`,
						},
					},
				},
			},
			funcMock: func(f fields) {
				f.logger.On("Info", "Processing 1 queue messages").Once()
				f.logger.On("Info", `#0 Message: {"dna_type":"dnaHuman","dna_sequence":"xxxx"}`)
				result := ResultMessage{}
				json.Unmarshal([]byte(`{"dna_type":"dnaHuman","dna_sequence":"xxxx"}`), &result)
				f.analysisRepository.On("GetAnalysisResult", result.DnaType, result.DnaSequence).Once().Return(domain.AnalysisModel{}, nil)
				f.analysisRepository.On("Register", commonAnalysis).Once().Return(nil)
				f.logger.On("Info", "Analysis result saved successfully").Once()
				f.statsRepository.On("NewDnaAnalysisWithName", result.DnaType).Once().Return(nil)
				f.logger.On("Info", "stats updated successfully").Once()
			},
			wantError: false,
			want:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.funcMock(tt.fields)
			handler := NewHandler(tt.fields.logger, tt.fields.analysisRepository, tt.fields.statsRepository)

			err := handler.Handle(tt.args.ctx, tt.args.req)
			if (err != nil) && tt.wantError {
				t.Fatal(domain.TestingErrorNotMatched(1, err, tt.want))
			}
		})
	}
}
