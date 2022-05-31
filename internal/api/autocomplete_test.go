package api

import (
	"net/http"
	"testing"

	"github.com/peopledatalabs/peopledatalabs-go/model"

	"github.com/stretchr/testify/assert"
)

func TestAutocomplete(t *testing.T) {
	// setup
	server := mockServer()
	defer server.Close()
	auto := Autocomplete{Client: mockClient(server)}

	// test
	params := model.AutocompleteParams{
		BaseParams:             model.BaseParams{Pretty: true, Size: 10},
		AutocompleteBaseParams: model.AutocompleteBaseParams{Field: "school", Text: "stanf"},
	}
	resp, err := auto.Autocomplete(params)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, resp.Status, http.StatusOK)
	assert.Equal(t, resp.Data[0].Name, "stanford university")
}
