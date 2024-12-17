package api

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/peopledatalabs/peopledatalabs-go/v3/model"

	"github.com/stretchr/testify/assert"
)

func TestAutocomplete(t *testing.T) {
	// setup
	auto := Autocomplete{Client: NewClient(os.Getenv("PDL_API_KEY"), "1.0.0")}

	// test
	params := model.AutocompleteParams{
		BaseParams:             model.BaseParams{Pretty: true, Size: 10},
		AutocompleteBaseParams: model.AutocompleteBaseParams{Field: "school", Text: "stanf"},
	}
	resp, err := auto.Autocomplete(context.Background(), params)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, resp.Status, http.StatusOK)
	assert.Equal(t, resp.Data[0].Name, "stanford university")
}

func TestAutocompleteClass(t *testing.T) {
	// setup
	auto := Autocomplete{Client: NewClient(os.Getenv("PDL_API_KEY"), "1.0.0")}

	// test
	params := model.AutocompleteParams{
		BaseParams:             model.BaseParams{Pretty: true, Size: 10, UpdatedTitleRoles: true},
		AutocompleteBaseParams: model.AutocompleteBaseParams{Field: "class", Text: "sale"},
	}
	resp, err := auto.Autocomplete(context.Background(), params)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, resp.Status, http.StatusOK)
	assert.Equal(t, resp.Data[0].Name, "sales_and_marketing")
}
