package api

import (
	"context"
	"fmt"

	"github.com/peopledatalabs/peopledatalabs-go/model"
)

const (
	personEnrichPath       = "/person/enrich"
	personIdentifyPath     = "/person/identify"
	personSearchPath       = "/person/search"
	personRetrievePath     = "/person/retrieve/%s"
	personBulkRetrievePath = "/person/retrieve/bulk"
)

type Person struct {
	Client
}

// Enrich a person record on a variety of fields
// TODO
func (p Person) Enrich(ctx context.Context, params model.EnrichPersonParams) (*model.EnrichPersonResponse, error) {
	var response *model.EnrichPersonResponse
	return response, p.Client.Get(ctx, personEnrichPath, params, response)
}

// Identify recovers all related profiles for an identity
// TODO
func (p Person) Identify(ctx context.Context, params model.IdentifyPersonParams) (*model.IdentifyPersonResponse, error) {
	var response *model.IdentifyPersonResponse
	return response, p.Client.Get(ctx, personIdentifyPath, params, response)
}

// Search
// TODO
func (p Person) Search(ctx context.Context, params model.SearchParams) (*model.SearchPersonResponse, error) {
	var response *model.SearchPersonResponse
	return response, p.Client.Post(ctx, personSearchPath, params, response)
}

// Retrieve recovers all related profiles for an identity
// TODO
func (p Person) Retrieve(ctx context.Context, params model.RetrievePersonParams) (*model.RetrievePersonResponse, error) {
	path := fmt.Sprintf(personRetrievePath, params.PersonID)
	var response *model.RetrievePersonResponse
	return response, p.Client.Get(ctx, path, params, &response)
}

// BulkRetrieve recovers all related profiles for an identity
// TODO
func (p Person) BulkRetrieve(ctx context.Context, params model.BulkRetrievePersonParams) ([]model.BulkRetrievePersonResponse, error) {
	var response []model.BulkRetrievePersonResponse
	return response, p.Client.Post(ctx, personBulkRetrievePath, params, &response)
}
