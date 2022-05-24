package peopledatalabs_go

import (
	"github.com/peopledatalabs/peopledatalabs-go/api"
)

const Version = "0.1.0"

type PLD struct {
	Person       api.Person
	Company      api.Company
	Location     api.Location
	School       api.School
	Autocomplete api.AutocompleteFunc
}

// TODO: Add docs
func New(apiKey string, opts ...api.ClientOptions) PLD {
	client := api.NewClient(apiKey, Version, opts...)

	return PLD{
		Person:       api.Person{Client: client},
		Company:      api.Company{Client: client},
		Location:     api.Location{Client: client},
		School:       api.School{Client: client},
		Autocomplete: api.Autocomplete{Client: client}.Autocomplete,
	}
}
