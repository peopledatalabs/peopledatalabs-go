package api

import (
	"context"

	"github.com/peopledatalabs/peopledatalabs-go/v6/model"
)

const (
	jobPostingSearchPath = "/job_posting/search"
)

type JobPosting struct {
	Client
}

// Search runs a search against the PDL job_posting dataset. Pass an
// Elasticsearch query via params.Query, or set any of the field-based filter
// fields on params for the parameter form. See:
// https://docs.peopledatalabs.com/docs/job-posting-search-api
func (j JobPosting) Search(ctx context.Context, params model.JobPostingSearchParams) (model.SearchJobPostingResponse, error) {
	if err := params.Validate(); err != nil {
		return model.SearchJobPostingResponse{}, err
	}
	var response model.SearchJobPostingResponse
	return response, j.Client.post(ctx, jobPostingSearchPath, params, &response)
}
