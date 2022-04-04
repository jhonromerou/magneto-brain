package awsapigateway

import (
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func Test_JSONStructResponse_Success(t *testing.T) {
	user := struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Name:     "Jhon Romero",
		Email:    "johnromero492@gmail.com",
		Password: "dificil-saberlo",
	}

	got := JSONStructResponse(http.StatusOK, user)

	expect := events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       `{"name":"Jhon Romero","email":"johnromero492@gmail.com","password":"dificil-saberlo"}`,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	assert.Equal(t, expect, got)
}

func Test_JSONStructResponse_Error(t *testing.T) {
	user := struct {
		Name interface{} `json:"name"`
	}{
		Name: make(chan int),
	}

	got := JSONStructResponse(http.StatusOK, user)

	expect := events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       `{"message":"json: unsupported type: chan int"}`,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	assert.Equal(t, expect, got)
}

func Test_ErrorAPIGatewayResponse_Error(t *testing.T) {
	err := ErrorAPIGatewayResponse{
		Message: "unexpected error",
	}

	got := err.Error()

	expect := "unexpected error"

	assert.Equal(t, expect, got)
}
