package api

import (
	"context"

	"github.com/peopledatalabs/peopledatalabs-go/model"
)

const (
	companyEnrichPath = "/company/enrich"
	companySearchPath = "/company/search"
	companyCleanPath  = "/company/clean"
)

type Company struct {
	Client
}

// Enrich a company
// docs: https://docs.peopledatalabs.com/docs/company-enrichment-api
func (c Company) Enrich(ctx context.Context, params model.EnrichCompanyParams) (model.EnrichCompanyResponse, error) {
	if err := params.Validate(); err != nil {
		return model.EnrichCompanyResponse{}, err
	}
	var response model.EnrichCompanyResponse
	return response, c.Client.get(ctx, companyEnrichPath, params, &response)
}

// Search gives you access to every record in our full Company dataset,
// which you can filter and segment using a search query.
// docs: https://docs.peopledatalabs.com/docs/company-search-api
func (c Company) Search(ctx context.Context, params model.SearchParams) (model.SearchCompanyResponse, error) {
	if err := params.Validate(); err != nil {
		return model.SearchCompanyResponse{}, err
	}
	var response model.SearchCompanyResponse
	return response, c.Client.post(ctx, companySearchPath, params, &response)
}

// Clean your company data, so you can better query our person data
// docs: https://docs.peopledatalabs.com/docs/cleaner-apis-reference#company-cleaner-api-companyclean-1
func (c Company) Clean(ctx context.Context, params model.CleanCompanyParams) (model.CleanCompanyResponse, error) {
	if err := params.Validate(); err != nil {
		return model.CleanCompanyResponse{}, err
	}
	var response model.CleanCompanyResponse
	return response, c.Client.get(ctx, companyCleanPath, params, &response)
}
