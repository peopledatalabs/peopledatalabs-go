package model

type CompanyParams struct {
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
	// TODO: Add validations of min params
}
type EnrichCompanyResponse struct {
	Status     int `json:"status"`
	Likelihood int `json:"likelihood"`
	Company
}

type CleanCompanyParams struct {
	BaseParams
	Name    string `json:"name,omitempty" url:"name,omitempty"`       // The name of the company
	Website string `json:"website,omitempty" url:"website,omitempty"` // A website the company uses
	Profile string `json:"profile,omitempty" url:"profile,omitempty"` // A social profile of the company (linkedin/facebook/twitter/crunchbase)
	// TODO: Add validations of min params

	//Ticker        string   `json:"ticker,omitempty" url:"ticker,omitempty"`                 // Company stock ticker
	//Location      []string `json:"location,omitempty" url:"location,omitempty"`             // Complete or partial company location
	//Locality      string   `json:"locality,omitempty" url:"locality,omitempty"`             //	Company locality
	//Region        string   `json:"region,omitempty" url:"region,omitempty"`                 //	Company region
	//Country       string   `json:"country,omitempty" url:"country,omitempty"`               //	Company country
	//StreetAddress string   `json:"street_address,omitempty" url:"street_address,omitempty"` //	Company address
	//PostalCode    string   `json:"postal_code,omitempty" url:"postal_code,omitempty"`       //	Company postal code

	// TODO: Check if we can use AdditionalParams (missing DataInclude, MinLikelihood, Required, IncludeIfMatched)
	// TODO: Check if we can use TitleCase
	//TitleCase        bool   `json:"titlecase" url:"titlecase,omitempty"`                   // Setting titlecase to true will titlecase the response data in 200 responses.
}
type CleanCompanyResponse struct {
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
