package api

import (
	"context"

	"github.com/peopledatalabs/peopledatalabs-go/model"
)

const (
	autocompletePath = "/Autocomplete"
)

type Autocomplete struct {
	Client
}

type AutocompleteFunc func(params model.AutocompleteParams) (model.AutocompleteResponse, error)
type AutocompleteWitchContextFunc func(ctx context.Context, params model.AutocompleteParams) (model.AutocompleteResponse, error)

// Autocomplete allows your users to get suggestions for Search API query values
// along with the number of available records for each suggestion.
// For example, schools starting with "stanf".
func (a Autocomplete) Autocomplete(params model.AutocompleteParams) (model.AutocompleteResponse, error) {
	return a.AutocompleteWithContext(context.Background(), params)
}
func (a Autocomplete) AutocompleteWithContext(ctx context.Context, params model.AutocompleteParams) (model.AutocompleteResponse, error) {
	if err := params.Validate(); err != nil {
		return model.AutocompleteResponse{}, err
	}
	var response model.AutocompleteResponse
	return response, a.Client.get(ctx, autocompletePath, params, &response)
}
