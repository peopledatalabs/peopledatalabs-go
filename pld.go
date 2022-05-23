package peopledatalabs_go

import (
	"github.com/peopledatalabs/peopledatalabs-go/api"
)

const Version = "0.1.0"

type PLD struct {
	Person  api.Person
	Company api.Company
}

// TODO: Add docs
func New(apiKey string, opts ...api.ClientOptions) PLD {
	client := api.NewClient(apiKey, Version, opts...)

	return PLD{
		Person:  api.Person{Client: client},
		Company: api.Company{Client: client},
	}
}
