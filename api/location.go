package api

import (
	"context"

	"github.com/peopledatalabs/peopledatalabs-go/v2/model"
)

const locationCleanPath = "/location/clean"

type Location struct {
	Client
}

// Clean your location data, so you can better query our person data
// docs: https://docs.peopledatalabs.com/docs/cleaner-apis-reference#location-cleaner-api-locationclean
func (l Location) Clean(ctx context.Context, params model.CleanLocationParams) (model.CleanLocationResponse, error) {
	if err := params.Validate(); err != nil {
		return model.CleanLocationResponse{}, err
	}
	var response model.CleanLocationResponse
	return response, l.Client.get(ctx, locationCleanPath, params, &response)
}
