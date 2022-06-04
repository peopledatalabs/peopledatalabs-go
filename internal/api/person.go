package api

import (
	"context"
	"fmt"

	"github.com/peopledatalabs/peopledatalabs-go/model"
)

const (
	personEnrichPath       = "/person/enrich"
	personBulkEnrichPath   = "/person/bulk"
	personIdentifyPath     = "/person/identify"
	personSearchPath       = "/person/search"
	personRetrievePath     = "/person/retrieve/%s"
	personBulkRetrievePath = "/person/retrieve/bulk"
)

type Person struct {
	Client
}

// Enrich a person record on a variety of fields
// docs: https://docs.peopledatalabs.com/docs/reference-person-enrichment-api
func (p Person) Enrich(params model.EnrichPersonParams) (model.EnrichPersonResponse, error) {
	return p.EnrichWithContext(context.Background(), params)
}
func (p Person) EnrichWithContext(ctx context.Context, params model.EnrichPersonParams) (model.EnrichPersonResponse, error) {
	if err := params.Validate(); err != nil {
		return model.EnrichPersonResponse{}, err
	}
	var response model.EnrichPersonResponse
	return response, p.Client.get(ctx, personEnrichPath, params, &response)
}

// BulkEnrich allows to enrich up to 100 persons in a single HTTP request
// docs: https://docs.peopledatalabs.com/docs/bulk-enrichment-api
func (p Person) BulkEnrich(params model.BulkEnrichPersonParams) ([]model.BulkEnrichPersonResponse, error) {
	return p.BulkEnrichWithContext(context.Background(), params)
}
func (p Person) BulkEnrichWithContext(ctx context.Context, params model.BulkEnrichPersonParams) ([]model.BulkEnrichPersonResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	var response []model.BulkEnrichPersonResponse
	return response, p.Client.post(ctx, personBulkEnrichPath, params, &response)
}

// Identify recovers all related profiles for an identity
// docs: https://docs.peopledatalabs.com/docs/identify-api
func (p Person) Identify(params model.IdentifyPersonParams) (model.IdentifyPersonResponse, error) {
	return p.IdentifyWithContext(context.Background(), params)
}
func (p Person) IdentifyWithContext(ctx context.Context, params model.IdentifyPersonParams) (model.IdentifyPersonResponse, error) {
	var response model.IdentifyPersonResponse
	return response, p.Client.get(ctx, personIdentifyPath, params, &response)
}

// Search is perfect for finding specific segments of people that you need to power your projects and products.
// It gives you direct access to our full API dataset.
// docs: https://docs.peopledatalabs.com/docs/person-search-api
func (p Person) Search(params model.SearchParams) (model.SearchPersonResponse, error) {
	return p.SearchWithContext(context.Background(), params)
}
func (p Person) SearchWithContext(ctx context.Context, params model.SearchParams) (model.SearchPersonResponse, error) {
	if err := params.Validate(); err != nil {
		return model.SearchPersonResponse{}, err
	}
	var response model.SearchPersonResponse
	return response, p.Client.post(ctx, personSearchPath, params, &response)
}

// Retrieve allows you to use a PDL Person ID to return data associated with that ID
// docs: https://docs.peopledatalabs.com/docs/person-retrieve-api
func (p Person) Retrieve(params model.RetrievePersonParams) (model.RetrievePersonResponse, error) {
	return p.RetrieveWithContext(context.Background(), params)
}
func (p Person) RetrieveWithContext(ctx context.Context, params model.RetrievePersonParams) (model.RetrievePersonResponse, error) {
	if err := params.Validate(); err != nil {
		return model.RetrievePersonResponse{}, err
	}
	path := fmt.Sprintf(personRetrievePath, params.PersonID)
	var response model.RetrievePersonResponse
	return response, p.Client.get(ctx, path, params, &response)
}

// BulkRetrieve allows to retrieve up to 100 persons in a single HTTP request
// docs: https://docs.peopledatalabs.com/docs/bulk-person-retrieve
func (p Person) BulkRetrieve(params model.BulkRetrievePersonParams) ([]model.BulkRetrievePersonResponse, error) {
	return p.BulkRetrieveWithContext(context.Background(), params)
}
func (p Person) BulkRetrieveWithContext(ctx context.Context, params model.BulkRetrievePersonParams) ([]model.BulkRetrievePersonResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	var response []model.BulkRetrievePersonResponse
	return response, p.Client.post(ctx, personBulkRetrievePath, params, &response)
}
