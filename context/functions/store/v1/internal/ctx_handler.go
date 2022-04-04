package internal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jhonromerou/magneto-brain/src/domain"

	"github.com/aws/aws-lambda-go/events"
)

// Handler defines entrypoint datatype for the lambda
type Handler struct {
	logger             domain.Logger
	analysisRepository domain.AnalysisRepository
	statsRepository    domain.StatsRepository
}

// Handle validates if dna sequence is of mutant or human.
func (h *Handler) Handle(_ context.Context, req events.SQSEvent) error {
	h.logger.Info(fmt.Sprintf("Processing %d queue messages", len(req.Records)))

	for i, message := range req.Records {
		h.logger.Info(fmt.Sprintf("#%d Message: %v", i, message.Body))

		result, err := NewResultMessageFromJSON(message.Body)
		if err != nil {
			h.logger.ErrorWithDetail(err, "cannot parse message")
			return nil
		}

		analysis, err := h.analysisRepository.GetAnalysisResult(result.DnaType, result.DnaSequence)
		if err != nil {
			h.logger.ErrorWithDetail(err, "error validate that analysis exists")
			return nil
		}

		if analysis.DnaType != "" {
			h.logger.Info("analysis already exist")
			return nil
		}

		item := domain.AnalysisModel{
			DnaType:     result.DnaType,
			DnaSequence: result.DnaSequence,
		}

		err = h.analysisRepository.Register(item)
		if err != nil {
			h.logger.ErrorWithDetail(err, "error saving analysis result")
			return nil
		}

		err = h.statsRepository.NewDnaAnalysisWithName(result.DnaType)
		if err != nil {
			h.logger.ErrorWithDetail(err, fmt.Sprintf("error updating stats of analysis %s", result.DnaType))
			return nil
		}

		h.logger.Info("stats updated successfully")
	}

	return nil
}

type ResultMessage struct {
	DnaType     string `json:"dna_type"`
	DnaSequence string `json:"dna_sequence"`
}

func (d *ResultMessage) isValid() error {
	if d.DnaType == "" {
		return errors.New("field dna_type is invalid or empty")
	}

	if d.DnaSequence == "" {
		return errors.New("field dna_sequence is invalid or empty")
	}

	return nil
}

func NewResultMessageFromJSON(body string) (ResultMessage, error) {
	message := ResultMessage{}

	err := json.Unmarshal([]byte(body), &message)
	if err != nil {
		return message, err
	}

	err = message.isValid()
	if err != nil {
		return message, err
	}

	return message, nil
}

// NewHandler creates a new instance of Handler
func NewHandler(logger domain.Logger, queueRepository domain.AnalysisRepository, statsRepository domain.StatsRepository) *Handler {
	return &Handler{
		logger:             logger,
		analysisRepository: queueRepository,
		statsRepository:    statsRepository,
	}
}
