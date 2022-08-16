package api

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/peopledatalabs/peopledatalabs-go/model"

	"github.com/stretchr/testify/assert"
)

func TestPerson_Enrich_Sandbox(t *testing.T) {
	// setup
	person := Person{Client: NewClient(os.Getenv("PDL_API_KEY"), "1.0.0")}
	person.Sandbox = true

	// test
	params := model.EnrichPersonParams{
		PersonParams:     model.PersonParams{Email: []string{"fletcherveronica@example.com"}},
		AdditionalParams: model.AdditionalParams{MinLikelihood: 6},
	}
	resp, err := person.Enrich(context.Background(), params)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.Status)
	assert.Equal(t, "twitter.com/omarmendez", resp.Data.TwitterUrl)
	assert.GreaterOrEqual(t, resp.Likelihood, 6)
}

func TestPerson_Identify_Sandbox(t *testing.T) {
	// setup
	person := Person{Client: NewClient(os.Getenv("PDL_API_KEY"), "1.0.0")}
	person.Sandbox = true

	// test
	params := model.IdentifyPersonParams{
		PersonParams: model.PersonParams{
			Company: []string{"walmart"},
		},
	}
	resp, err := person.Identify(context.Background(), params)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.Status)
	if assert.Greater(t, len(resp.Matches), 1) {
		assert.Greater(t, resp.Matches[0].MatchScore, resp.Matches[1].MatchScore)
	}
}

func TestPerson_Search_Sandbox(t *testing.T) {
	// setup
	person := Person{Client: NewClient(os.Getenv("PDL_API_KEY"), "1.0.0")}
	person.Sandbox = true
	numResults := 3

	// test
	params := model.SearchParams{
		BaseParams: model.BaseParams{Size: numResults},
		SearchBaseParams: model.SearchBaseParams{
			SQL: "SELECT * FROM person WHERE location_country='mexico';",
		},
	}
	resp, err := person.Search(context.Background(), params)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.Status)
	assert.Equal(t, numResults, len(resp.Data))
	assert.NotEmpty(t, resp.ScrollToken)
}
