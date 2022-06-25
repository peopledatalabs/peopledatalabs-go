package api

import (
	"context"

	"github.com/peopledatalabs/peopledatalabs-go/model"
)

const (
	autocompletePath = "/autocomplete"
)

type Autocomplete struct {
	Client
}

type AutocompleteFunc func(ctx context.Context, params model.AutocompleteParams) (model.AutocompleteResponse, error)

// Autocomplete allows your users to get suggestions for Search API query values
// along with the number of available records for each suggestion.
// For example, schools starting with "stanf".
func (a Autocomplete) Autocomplete(ctx context.Context, params model.AutocompleteParams) (model.AutocompleteResponse, error) {
	if err := params.Validate(); err != nil {
		return model.AutocompleteResponse{}, err
	}
	var response model.AutocompleteResponse
	return response, a.Client.get(ctx, autocompletePath, params, &response)
}
