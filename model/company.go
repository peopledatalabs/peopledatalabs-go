package model

import "errors"

type CompanyParams struct {
	Id            []string `json:"pdl_id,omitempty" url:"pdl_id,omitempty"`                 // The pdl_id of a company
	Name          string   `json:"name,omitempty" url:"name,omitempty"`                     // The name of the company
	Website       string   `json:"website,omitempty" url:"website,omitempty"`               // A website the company uses
	Profile       string   `json:"profile,omitempty" url:"profile,omitempty"`               // A social profile of the company (linkedin/facebook/twitter/crunchbase)
	Ticker        string   `json:"ticker,omitempty" url:"ticker,omitempty"`                 // Company stock ticker
	Location      []string `json:"location,omitempty" url:"location,omitempty"`             // Complete or partial company location
	Locality      string   `json:"locality,omitempty" url:"locality,omitempty"`             //	Company locality
	Region        string   `json:"region,omitempty" url:"region,omitempty"`                 //	Company region
	Country       string   `json:"country,omitempty" url:"country,omitempty"`               //	Company country
	StreetAddress string   `json:"street_address,omitempty" url:"street_address,omitempty"` //	Company address
	PostalCode    string   `json:"postal_code,omitempty" url:"postal_code,omitempty"`       //	Company postal code
}

type EnrichCompanyParams struct {
	BaseParams
	CompanyParams
	AdditionalParams
}

func (params EnrichCompanyParams) Validate() error {
	if params.Name == "" && params.Ticker == "" && params.Website == "" && params.Profile == "" {
		return errors.New("enrich company: you must input a name OR ticker OR website OR profile")
	}
	return nil
}

type EnrichCompanyResponse struct {
	Status     int `json:"status"`
	Likelihood int `json:"likelihood"`
	Company
}

type CleanCompanyParams struct {
	Name    string `json:"name,omitempty" url:"name,omitempty"`       // The name of the company
	Website string `json:"website,omitempty" url:"website,omitempty"` // A website the company uses
	Profile string `json:"profile,omitempty" url:"profile,omitempty"` // A social profile of the company (linkedin/facebook/twitter/crunchbase)
}

func (params CleanCompanyParams) Validate() error {
	if params.Name == "" && params.Website == "" && params.Profile == "" {
		return errors.New("clean company: you must input a name OR website OR profile")
	}
	return nil
}

type CleanCompanyResponse struct {
	Status int `json:"status"`
	Company
	FuzzyMatch bool `json:"fuzzy_match"`
}

type SearchCompanyResponse struct {
	Status      int       `json:"status"`
	Data        []Company `json:"data"`
	ScrollToken string    `json:"scroll_token"` // Scroll value used for pagination
	Total       int       `json:"total"`        // Number of records matching a given query or sql input.
	Error       struct {
		Type    []string `json:"type"`
		Message string   `json:"message"`
	} `json:"error"`
}
