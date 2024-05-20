package api

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/peopledatalabs/peopledatalabs-go/v2/model"

	"github.com/stretchr/testify/assert"
)

func TestPerson_Enrich_Sandbox(t *testing.T) {
	// setup
	client := NewClient(os.Getenv("PDL_API_KEY"), "1.0.0", ClientOptions(func(c *Client) {
		c.Sandbox = true
	}))
	person := Person{Client: client}

	// test
	params := model.EnrichPersonParams{
		PersonParams:     model.PersonParams{Email: []string{"reneewillis74@aol.com"}},
		AdditionalParams: model.AdditionalParams{MinLikelihood: 6},
	}
	resp, err := person.Enrich(context.Background(), params)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.Status)
	assert.Equal(t, "twitter.com/rc1994", resp.Data.TwitterUrl)
	assert.GreaterOrEqual(t, resp.Likelihood, 6)
}

func TestPerson_Identify_Sandbox(t *testing.T) {
	// setup
	client := NewClient(os.Getenv("PDL_API_KEY"), "1.0.0", ClientOptions(func(c *Client) {
		c.Sandbox = true
	}))
	person := Person{Client: client}

	// test
	params := model.IdentifyPersonParams{
		PersonParams: model.PersonParams{
			Company: []string{"adams group"},
		},
	}
	resp, err := person.Identify(context.Background(), params)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.Status)
	if assert.Greater(t, len(resp.Matches), 1) {
		assert.GreaterOrEqual(t, resp.Matches[0].MatchScore, resp.Matches[1].MatchScore)
	}
}

func TestPerson_Search_Sandbox(t *testing.T) {
	// setup
	client := NewClient(os.Getenv("PDL_API_KEY"), "1.0.0", ClientOptions(func(c *Client) {
		c.Sandbox = true
	}))
	person := Person{Client: client}
	numResults := 3

	// test
	params := model.SearchParams{
		BaseParams: model.BaseParams{Size: numResults},
		SearchBaseParams: model.SearchBaseParams{
			SQL: "SELECT * FROM person WHERE location_country='united states';",
		},
	}
	resp, err := person.Search(context.Background(), params)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.Status)
	assert.Equal(t, numResults, len(resp.Data))
	assert.NotEmpty(t, resp.ScrollToken)
}
