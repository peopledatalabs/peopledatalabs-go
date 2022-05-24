package api

import (
	"context"

	"github.com/peopledatalabs/peopledatalabs-go/model"
)

const locationCleanPath = "/location/clean"

type Location struct {
	Client
}

func (l Location) Clean(ctx context.Context, params model.CleanLocationParams) (model.CleanLocationResponse, error) {
	var response model.CleanLocationResponse
	return response, l.Client.Get(ctx, locationCleanPath, params, &response)
}
