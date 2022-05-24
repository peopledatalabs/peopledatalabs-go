package api

import (
	"context"

	"github.com/peopledatalabs/peopledatalabs-go/model"
)

const (
	companyEnrichPath = "/company/enrich"
	companySearchPath = "/company/search"
	companyCleanPath  = "/company/enrich"
)

type Company struct {
	Client
}

// Enrich a company
// TODO: docs
func (c Company) Enrich(ctx context.Context, params model.EnrichCompanyParams) (model.EnrichCompanyResponse, error) {
	var response model.EnrichCompanyResponse
	return response, c.Client.Get(ctx, companyEnrichPath, params, &response)
}

// Search
// TODO: docs
func (c Company) Search(ctx context.Context, params model.SearchParams) (model.SearchCompanyResponse, error) {
	var response model.SearchCompanyResponse
	return response, c.Client.Post(ctx, companySearchPath, params, &response)
}

// Clean
// TODO: docs
func (c Company) Clean(ctx context.Context, params model.CleanCompanyParams) (model.CleanCompanyResponse, error) {
	var response model.CleanCompanyResponse
	return response, c.Client.Get(ctx, companyCleanPath, params, &response)
}
