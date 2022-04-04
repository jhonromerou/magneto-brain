package internal

import (
	"context"
	"fmt"
	"math"
	"net/http"

	"github.com/jhonromerou/magneto-brain/src/domain"
	"github.com/jhonromerou/magneto-brain/src/infrastructure/awsapigateway"

	"github.com/aws/aws-lambda-go/events"
)

type StatsResponse struct {
	Dna StatsResponseCounts `json:"dna"`
}

type StatsResponseCounts struct {
	Mutants uint    `json:"count_mutant_dna"`
	Humans  uint    `json:"count_human_dna"`
	Ratio   float64 `json:"ratio"`
}

// Handler defines entrypoint datatype for the lambda
type Handler struct {
	logger          domain.Logger
	statsRepository domain.StatsRepository
}

// Handle gets stats of analysis dna.
func (h *Handler) Handle(_ context.Context, events events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	h.logger.Info("Getting stats")
	humans, err := h.statsRepository.GetDnaAnalysisByName(domain.STATS_NAME_DNA_HUMAN)

	if err != nil {
		h.logger.ErrorWithDetail(err, fmt.Sprintf("error getting %s stats", domain.STATS_NAME_DNA_HUMAN))
		return awsapigateway.JSONStructResponse(http.StatusInternalServerError, err), nil
	}

	mutants, err := h.statsRepository.GetDnaAnalysisByName(domain.STATS_NAME_DNA_MUTANT)
	if err != nil {
		h.logger.ErrorWithDetail(err, fmt.Sprintf("error getting %s stats", domain.STATS_NAME_DNA_MUTANT))
		return awsapigateway.JSONStructResponse(http.StatusInternalServerError, err), nil
	}
	totalAnalysis := humans.Quantity + mutants.Quantity
	var ratio float64
	if totalAnalysis > 0 {
		ratio = float64(mutants.Quantity) / float64(totalAnalysis)
		ratio = math.Floor(ratio*100) / 100
	}

	response := StatsResponse{
		StatsResponseCounts{
			Mutants: uint(mutants.Quantity),
			Humans:  uint(humans.Quantity),
			Ratio:   ratio,
		},
	}
	return awsapigateway.JSONStructResponse(http.StatusOK, response), nil
}

// NewHandler creates a new instance of Handler
func NewHandler(logger domain.Logger, statsRepository domain.StatsRepository) *Handler {
	return &Handler{
		logger:          logger,
		statsRepository: statsRepository,
	}
}
