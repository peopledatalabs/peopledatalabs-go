package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/peopledatalabs/peopledatalabs-go/model"

	"github.com/stretchr/testify/assert"
)

func TestClient_RestError(t *testing.T) {
	params := model.EnrichPersonParams{
		PersonParams: model.PersonParams{Name: []string{"not existing name"}},
	}
	err := NewClient(os.Getenv("PDL_API_KEY"), "1.0.0").get(context.Background(), "/person/enrich", params, nil)

	assert.Error(t, err)
	if assert.IsType(t, model.RestError{}, err) {
		restError := err.(model.RestError)
		assert.Equal(t, http.StatusBadRequest, restError.Status)
	}
}

func TestClient_NotFound(t *testing.T) {
	params := model.CleanCompanyParams{
		Name: "not existing company",
	}
	err := NewClient(os.Getenv("PDL_API_KEY"), "1.0.0").get(context.Background(), "/company/clean", params, nil)

	assert.Error(t, err)
	if assert.IsType(t, model.NotFoundError{}, err) {
		nfError := err.(model.NotFoundError)
		assert.Equal(t, http.StatusNotFound, nfError.Status)
	}
}

func TestClient_stdError(t *testing.T) {
	// setup
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotImplemented)
			w.Write([]byte(`endpoint not yet implemented`))
		}))
	defer server.Close()

	client := NewClient(os.Getenv("PDL_API_KEY"), "1.0.0")
	client.BaseURL = server.URL
	client.ApiVersion = ""

	// test
	params := model.EnrichPersonParams{
		PersonParams: model.PersonParams{Name: []string{"sean"}},
	}
	err := client.get(context.Background(), "/new-endpoint", params, nil)

	// assertions
	assert.Error(t, err)
	_, isRestErrorType := err.(model.RestError)
	assert.False(t, isRestErrorType)
}

func TestClient_ErrorUnmarshall(t *testing.T) {
	// setup
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"invalid_json": something really wrong happened on PDL server`))
		}))
	defer server.Close()

	client := NewClient(os.Getenv("PDL_API_KEY"), "1.0.0")
	client.BaseURL = server.URL
	client.ApiVersion = ""

	// test
	params := model.EnrichPersonParams{
		PersonParams: model.PersonParams{Name: []string{"sean"}},
	}
	var out model.EnrichPersonResponse
	err := client.get(context.Background(), "/error-json-response", params, &out)

	// assertion
	assert.True(t, strings.HasPrefix(err.Error(), "makeRequest: error while reading response body"))
}
