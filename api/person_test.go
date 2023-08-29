package api

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/peopledatalabs/peopledatalabs-go/model"

	"github.com/stretchr/testify/assert"
)

func TestPerson_Enrich(t *testing.T) {
	// setup
	person := Person{Client: NewClient(os.Getenv("PDL_API_KEY"), "1.0.0")}

	// test
	params := model.EnrichPersonParams{
		PersonParams:     model.PersonParams{Profile: []string{"http://linkedin.com/in/seanthorne"}},
		AdditionalParams: model.AdditionalParams{MinLikelihood: 6},
	}
	resp, err := person.Enrich(context.Background(), params)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.Status)
	assert.Equal(t, "twitter.com/seanthorne5", resp.Data.TwitterUrl)
	assert.GreaterOrEqual(t, resp.Likelihood, 6)
}

func TestPerson_BulkEnrich(t *testing.T) {
	// setup
	person := Person{Client: NewClient(os.Getenv("PDL_API_KEY"), "1.0.0")}

	// test
	params := model.BulkEnrichPersonParams{
		BaseParams:       model.BaseParams{Pretty: true},
		AdditionalParams: model.AdditionalParams{MinLikelihood: 6, IncludeIfMatched: true, Required: "full_name"},
		Requests: []model.BulkEnrichSinglePersonParams{
			{
				Params: model.PersonParams{
					Profile: []string{"linkedin.com/in/seanthorne"},
				},
			},
			{
				Params: model.PersonParams{
					Profile: []string{"https://www.linkedin.com/in/haydenconrad/"},
				},
			},
		},
	}
	resp, err := person.BulkEnrich(context.Background(), params)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp[0].Status)
	assert.Equal(t, http.StatusOK, resp[1].Status)
	assert.Equal(t, "sean thorne", resp[0].Data.FullName)
	assert.Equal(t, "hayden conrad", resp[1].Data.FullName)
}

func TestPerson_Identify(t *testing.T) {
	// setup
	person := Person{Client: NewClient(os.Getenv("PDL_API_KEY"), "1.0.0")}

	// test
	params := model.IdentifyPersonParams{
		PersonParams: model.PersonParams{
			FirstName: []string{"sean"},
			LastName:  []string{"thorne"},
			Company:   []string{"people data labs"},
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

func TestPerson_Search(t *testing.T) {
	// setup
	person := Person{Client: NewClient(os.Getenv("PDL_API_KEY"), "1.0.0")}
	numResults := 3

	// test
	params := model.SearchParams{
		BaseParams: model.BaseParams{Size: numResults},
		SearchBaseParams: model.SearchBaseParams{
			SQL: "SELECT * FROM person WHERE location_country='mexico' AND job_title_role='health' AND phone_numbers IS NOT NULL;",
		},
	}
	resp, err := person.Search(context.Background(), params)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.Status)
	assert.Equal(t, numResults, len(resp.Data))
	assert.NotEmpty(t, resp.ScrollToken)
}

func TestPerson_Retrieve(t *testing.T) {
	// setup
	person := Person{Client: NewClient(os.Getenv("PDL_API_KEY"), "1.0.0")}
	id := "qEnOZ5Oh0poWnQ1luFBfVw_0000"

	// test
	params := model.RetrievePersonParams{PersonID: id}
	resp, err := person.Retrieve(context.Background(), params)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.Status)
	assert.Equal(t, id, resp.Data.Id)
	assert.Equal(t, "sean thorne", resp.Data.FullName)
}

func TestPerson_BulkRetrieve(t *testing.T) {
	// setup
	person := Person{Client: NewClient(os.Getenv("PDL_API_KEY"), "1.0.0")}
	id1 := "qEnOZ5Oh0poWnQ1luFBfVw_0000"
	id2 := "9Grd31hT3RFKVzsyecBGPg_0000"

	// test
	params := model.BulkRetrievePersonParams{
		Requests: []model.BulkRetrieveSinglePersonParams{
			{ID: id1},
			{ID: id2},
		},
	}
	resp, err := person.BulkRetrieve(context.Background(), params)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, len(params.Requests), len(resp))
	assert.Equal(t, http.StatusOK, resp[0].Status)
	assert.Equal(t, "sean thorne", resp[0].Data.FullName)
	assert.Equal(t, http.StatusOK, resp[1].Status)
	assert.Equal(t, "varun villait", resp[1].Data.FullName)
}
