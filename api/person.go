package api

import (
	"context"
	"fmt"

	"github.com/peopledatalabs/peopledatalabs-go/v6/model"
)

const (
	personEnrichPath       = "/person/enrich"
	personBulkEnrichPath   = "/person/bulk"
	personIdentifyPath     = "/person/identify"
	personSearchPath       = "/person/search"
	personRetrievePath     = "/person/retrieve/%s"
	personBulkRetrievePath = "/person/retrieve/bulk"
	personChangelogPath    = "/person/changelog"
)

type Person struct {
	Client
}

// Enrich a person record on a variety of fields
// docs: https://docs.peopledatalabs.com/docs/reference-person-enrichment-api
func (p Person) Enrich(ctx context.Context, params model.EnrichPersonParams) (model.EnrichPersonResponse, error) {
	if err := params.Validate(); err != nil {
		return model.EnrichPersonResponse{}, err
	}
	var response model.EnrichPersonResponse
	return response, p.Client.get(ctx, personEnrichPath, params, &response)
}

// BulkEnrich allows to enrich up to 100 persons in a single HTTP request
// docs: https://docs.peopledatalabs.com/docs/bulk-enrichment-api
func (p Person) BulkEnrich(ctx context.Context, params model.BulkEnrichPersonParams) ([]model.BulkEnrichPersonResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	var response []model.BulkEnrichPersonResponse
	return response, p.Client.post(ctx, personBulkEnrichPath, params, &response)
}

// Identify recovers all related profiles for an identity
// docs: https://docs.peopledatalabs.com/docs/identify-api
func (p Person) Identify(ctx context.Context, params model.IdentifyPersonParams) (model.IdentifyPersonResponse, error) {
	var response model.IdentifyPersonResponse
	return response, p.Client.get(ctx, personIdentifyPath, params, &response)
}

// Search is perfect for finding specific segments of people that you need to power your projects and products.
// It gives you direct access to our full API dataset.
// docs: https://docs.peopledatalabs.com/docs/person-search-api
func (p Person) Search(ctx context.Context, params model.SearchParams) (model.SearchPersonResponse, error) {
	if err := params.Validate(); err != nil {
		return model.SearchPersonResponse{}, err
	}
	var response model.SearchPersonResponse
	return response, p.Client.post(ctx, personSearchPath, params, &response)
}

// Retrieve allows you to use a PDL Person ID to return data associated with that ID
// docs: https://docs.peopledatalabs.com/docs/person-retrieve-api
func (p Person) Retrieve(ctx context.Context, params model.RetrievePersonParams) (model.RetrievePersonResponse, error) {
	if err := params.Validate(); err != nil {
		return model.RetrievePersonResponse{}, err
	}
	path := fmt.Sprintf(personRetrievePath, params.PersonID)
	var response model.RetrievePersonResponse
	return response, p.Client.get(ctx, path, params, &response)
}

// BulkRetrieve allows to retrieve up to 100 persons in a single HTTP request
// docs: https://docs.peopledatalabs.com/docs/bulk-person-retrieve
func (p Person) BulkRetrieve(ctx context.Context, params model.BulkRetrievePersonParams) ([]model.BulkRetrievePersonResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	var response []model.BulkRetrievePersonResponse
	return response, p.Client.post(ctx, personBulkRetrievePath, params, &response)
}

// Changelog returns a list of changes made to the person data
// docs: https://docs.peopledatalabs.com/docs/person-changelog-api
func (p Person) Changelog(ctx context.Context, params model.ChangelogPersonParams) (model.ChangelogPersonResponse, error) {
	if err := params.Validate(); err != nil {
		return model.ChangelogPersonResponse{}, err
	}
	var response model.ChangelogPersonResponse
	return response, p.Client.post(ctx, personChangelogPath, params, &response)
}
