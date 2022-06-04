package api

import (
	"context"

	"github.com/peopledatalabs/peopledatalabs-go/model"
)

const schoolCleanPath = "/school/clean"

type School struct {
	Client
}

// Clean your school data, so you can better query our person data
// docs: https://docs.peopledatalabs.com/docs/cleaner-apis-reference#school-cleaner-api-schoolclean
func (s School) Clean(params model.CleanSchoolParams) (model.CleanSchoolResponse, error) {
	return s.CleanWithContext(context.Background(), params)
}
func (s School) CleanWithContext(ctx context.Context, params model.CleanSchoolParams) (model.CleanSchoolResponse, error) {
	if err := params.Validate(); err != nil {
		return model.CleanSchoolResponse{}, err
	}
	var response model.CleanSchoolResponse
	return response, s.Client.get(ctx, schoolCleanPath, params, &response)
}
