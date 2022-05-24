package api

import (
	"context"

	"github.com/peopledatalabs/peopledatalabs-go/model"
)

const schoolCleanPath = "/school/clean"

type School struct {
	Client
}

func (s School) Clean(ctx context.Context, params model.CleanSchoolParams) (model.CleanSchoolResponse, error) {
	var response model.CleanSchoolResponse
	return response, s.Client.Get(ctx, schoolCleanPath, params, &response)
}
