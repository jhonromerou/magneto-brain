package awsapigateway

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// ErrorAPIGatewayResponse defines a standard error for http response
type ErrorAPIGatewayResponse struct {
	ErrorCode string `json:"error,omitempty"`
	Message   string `json:"message,omitempty"`
}

// Error prints the default error message
func (e *ErrorAPIGatewayResponse) Error() string {
	return e.Message
}

// JSONError create an error response from a given status code and error
func JSONError(statusCode int, err error) events.APIGatewayProxyResponse {
	body, _ := json.Marshal(ErrorAPIGatewayResponse{
		Message: err.Error(),
	})

	return JSONResponse(statusCode, body)
}

// JSONResponse create response from a given status code and body
func JSONResponse(statusCode int, body []byte) events.APIGatewayProxyResponse {
	buf := bytes.Buffer{}
	json.HTMLEscape(&buf, body)

	return events.APIGatewayProxyResponse{
		Body:       buf.String(),
		StatusCode: statusCode,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}

// JSONStructResponse create an error response from a given status code
// and string body
func JSONStructResponse(statusCode int, structure interface{}) events.APIGatewayProxyResponse {
	raw, err := json.Marshal(structure)
	if err != nil {
		return JSONError(http.StatusInternalServerError, err)
	}

	return JSONResponse(statusCode, raw)
}
