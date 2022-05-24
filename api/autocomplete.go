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

func (a Autocomplete) Autocomplete(ctx context.Context, params model.AutocompleteParams) (model.AutocompleteResponse, error) {
	var response model.AutocompleteResponse
	return response, a.Client.Get(ctx, autocompletePath, params, &response)
}
