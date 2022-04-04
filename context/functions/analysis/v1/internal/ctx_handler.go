package internal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jhonromerou/magneto-brain/src/domain"
	"github.com/jhonromerou/magneto-brain/src/domain/services"
	"github.com/jhonromerou/magneto-brain/src/infrastructure/awsapigateway"

	"github.com/aws/aws-lambda-go/events"
)

// Handler defines entrypoint datatype for the lambda
type Handler struct {
	logger          domain.Logger
	queueRepository domain.QueueAnalysisRepository
	dnaPlanSequence string
}

// Handle analyzes dna sequence type.
func (h *Handler) Handle(_ context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	h.logger.Info(fmt.Sprintf("Processing analysis with data %s", req.Body))

	dna, err := services.NewDnaFromBody(req.Body)
	if err != nil {
		return h.errorWithDnaPattern(err), nil
	}

	h.dnaPlanSequence = dna.ToString()
	dnaTableSequence := dna.GetTable()
	if dnaTableSequence == nil {
		return h.dnaPatterIsInvalid(), nil
	}

	mutantService := services.NewMutantService(dnaTableSequence)
	isMutant := mutantService.IsMutant()
	dnaType := domain.STATS_NAME_DNA_HUMAN
	if isMutant {
		dnaType = domain.STATS_NAME_DNA_MUTANT
	}

	errSaving := h.saveValidationDna(dnaType)
	if errSaving.Name != "" {
		return awsapigateway.JSONStructResponse(http.StatusInternalServerError, errSaving), nil
	}

	if isMutant {
		return h.dnaTypeIsOfMutant(), nil
	}

	return h.dnaTypeIsOfHuman(), nil
}

func (h *Handler) saveValidationDna(dnaType string) domain.AnalysisValidatorResponse {
	message := fmt.Sprintf(`{"dna_type":"%s","dna_sequence":"%s"}`, dnaType, h.dnaPlanSequence)
	errSaving := h.queueRepository.PublishMessage(message)
	if errSaving != nil {
		h.logger.ErrorWithDetail(errSaving, "error publishing message on the queue")
		return domain.AnalysisValidatorResponse{
			Name:   "dna_error",
			Format: "error publishing message on the queue",
		}
	}
	return domain.AnalysisValidatorResponse{}
}

func (h *Handler) dnaTypeIsOfMutant() events.APIGatewayProxyResponse {
	response := domain.AnalysisValidatorResponse{
		Name:   "dna_is_of_mutant",
		Format: "request with dna {dna} is of mutant",
		Values: map[string]string{
			"dna": h.dnaPlanSequence,
		},
	}

	return awsapigateway.JSONStructResponse(http.StatusOK, response)
}

func (h *Handler) dnaTypeIsOfHuman() events.APIGatewayProxyResponse {
	response := domain.AnalysisValidatorResponse{
		Name:   "dna_not_is_of_mutant",
		Format: "request with dna {dna} not is of mutant",
		Values: map[string]string{
			"dna": h.dnaPlanSequence,
		},
	}
	return awsapigateway.JSONStructResponse(http.StatusForbidden, response)
}

func (h *Handler) errorWithDnaPattern(err error) events.APIGatewayProxyResponse {
	h.logger.ErrorWithDetail(err, "error with dna pattern sequence")

	response := domain.AnalysisValidatorResponse{
		Name:   "dna_error",
		Format: "error with dna {dna} pattern sequence",
		Values: map[string]string{
			"dna": h.dnaPlanSequence,
		},
	}

	return awsapigateway.JSONStructResponse(http.StatusBadRequest, response)
}

func (h *Handler) dnaPatterIsInvalid() events.APIGatewayProxyResponse {
	h.logger.Error("dna pattern sequence invalid")

	err := domain.AnalysisValidatorResponse{
		Name:   "dna_patter_sequence_invalid",
		Format: "error validating that dna {dna} have allow pattern sequence",
		Values: map[string]string{
			"dna": h.dnaPlanSequence,
		},
	}
	return awsapigateway.JSONStructResponse(http.StatusBadRequest, err)
}

// NewHandler creates a new instance of Handler.
func NewHandler(logger domain.Logger, queueRepository domain.QueueAnalysisRepository) *Handler {
	return &Handler{
		logger:          logger,
		queueRepository: queueRepository,
	}
}
