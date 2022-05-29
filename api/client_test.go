package api

import (
	"context"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/peopledatalabs/peopledatalabs-go/model"

	"github.com/stretchr/testify/assert"
)

func TestClient_RestError(t *testing.T) {
	server := customMockServer(http.StatusNotFound, `{"status":404,"error":{"type":["not_found"],"message":"No records were found matching your request"}}`)
	defer server.Close()
	client := mockClient(server)

	params := model.EnrichPersonParams{
		PersonParams: model.PersonParams{Name: []string{"not existing name"}},
	}
	err := client.Get(context.Background(), "/person/enrich", params, nil)

	assert.Error(t, err)
	if assert.IsType(t, model.RestError{}, err) {
		restError := err.(model.RestError)
		assert.Equal(t, http.StatusNotFound, restError.Status)
	}
}

func TestClient_GenericError(t *testing.T) {
	server := customMockServer(http.StatusNotImplemented, `endpoint not yet implemented`)
	defer server.Close()
	client := mockClient(server)

	params := model.EnrichPersonParams{
		PersonParams: model.PersonParams{Name: []string{"sean"}},
	}
	err := client.Get(context.Background(), "/new-endpoint", params, nil)

	assert.Error(t, err)
	_, isRestErrorType := err.(model.RestError)
	assert.False(t, isRestErrorType)
}

func TestClient_ErrorUnmarshall(t *testing.T) {
	server := customMockServer(http.StatusOK, `{"invalid_json": aaaaaaa}`)
	defer server.Close()
	client := mockClient(server)

	params := model.EnrichPersonParams{
		PersonParams: model.PersonParams{Name: []string{"sean"}},
	}
	var out model.EnrichPersonResponse
	err := client.Get(context.Background(), "/error-json-response", params, &out)

	assert.True(t, strings.HasPrefix(err.Error(), "makeRequest: error while reading response body"))
}

func TestNewClient_CustomParams(t *testing.T) {
	var options []ClientOptions
	options = append(options, WithTimeout(30*time.Second))
	options = append(options, WithAPIVersion("v5"))
	options = append(options, WithHTTPClient(http.DefaultClient))
	_ = NewClient("api_key", "1", options...)
}
