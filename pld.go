package peopledatalabs_go

import (
	"github.com/peopledatalabs/peopledatalabs-go/internal/api"
)

const Version = "1.0.1"

type pld struct {
	Person                  api.Person
	Company                 api.Company
	Location                api.Location
	School                  api.School
	Autocomplete            api.AutocompleteFunc
	AutocompleteWithContext api.AutocompleteWitchContextFunc
}

// New returns a new People Data Labs Client
// All interactions with the API must be done using this client
//
// You must provide an API KEY. Get your API key by creating a free account at https://www.peopledatalabs.com/signup.
// You can provide api.ClientOptions to customize client behaviour
func New(apiKey string, opts ...api.ClientOptions) *pld {
	client := api.NewClient(apiKey, Version, opts...)

	autocompleteClient := api.Autocomplete{Client: client}
	return &pld{
		Person:                  api.Person{Client: client},
		Company:                 api.Company{Client: client},
		Location:                api.Location{Client: client},
		School:                  api.School{Client: client},
		Autocomplete:            autocompleteClient.Autocomplete,
		AutocompleteWithContext: autocompleteClient.AutocompleteWithContext,
	}
}
