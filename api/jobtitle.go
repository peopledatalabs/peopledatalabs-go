package api

import (
	"context"

	"github.com/peopledatalabs/peopledatalabs-go/v2/model"
)

const (
	jobTitlePath = "/job_title/enrich"
)

type JobTitle struct {
	Client
}

type JobTitleFunc func(ctx context.Context, params model.JobTitleParams) (model.JobTitleResponse, error)

// JobTitle allows your users to get information for JobTitle values
func (a JobTitle) JobTitle(ctx context.Context, params model.JobTitleParams) (model.JobTitleResponse, error) {
	if err := params.Validate(); err != nil {
		return model.JobTitleResponse{}, err
	}
	var response model.JobTitleResponse
	return response, a.Client.get(ctx, jobTitlePath, params, &response)
}
