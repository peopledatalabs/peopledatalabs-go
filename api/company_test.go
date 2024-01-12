package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/peopledatalabs/peopledatalabs-go/model"

	"github.com/stretchr/testify/assert"
)

func TestCompany_Enrich(t *testing.T) {
	// setup
	company := Company{Client: NewClient(os.Getenv("PDL_API_KEY"), "1.0.0")}

	// test
	params := model.EnrichCompanyParams{
		CompanyParams: model.CompanyParams{Name: "Google"},
	}
	resp, err := company.Enrich(context.Background(), params)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.Status)
	assert.Equal(t, 1998, resp.Founded)
}

func TestCompany_BulkEnrich(t *testing.T) {
	// setup
	company := Company{Client: NewClient(os.Getenv("PDL_API_KEY"), "1.0.0")}

	// test
	params := model.BulkEnrichCompanyParams{
		BaseParams: model.BaseParams{Pretty: true},
		Requests: []model.BulkEnrichSingleCompanyParams{
			{
				Params: model.CompanyParams{
					Profile: "linkedin.com/company/peopledatalabs/",
				},
			},
			{
				Params: model.CompanyParams{
					Profile: "https://www.linkedin.com/company/apple/",
				},
			},
		},
	}
	resp, err := company.BulkEnrich(context.Background(), params)

	fmt.Println(resp)
	// assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp[0].Status)
	assert.Equal(t, http.StatusOK, resp[1].Status)
	assert.Equal(t, "People Data Labs", resp[0].DisplayName)
	assert.Equal(t, "Apple", resp[1].DisplayName)
}

func TestCompany_Clean(t *testing.T) {
	// setup
	company := Company{Client: NewClient(os.Getenv("PDL_API_KEY"), "1.0.0")}

	// test
	params := model.CleanCompanyParams{Website: "apple.com"}
	resp, err := company.Clean(context.Background(), params)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, "apple", resp.Name)
}

func TestCompany_Search(t *testing.T) {
	// setup
	company := Company{Client: NewClient(os.Getenv("PDL_API_KEY"), "1.0.0")}

	// test
	params := model.SearchParams{
		BaseParams:       model.BaseParams{Size: 100},
		SearchBaseParams: model.SearchBaseParams{SQL: "SELECT * FROM company WHERE name='people data labs'"},
	}
	resp, err := company.Search(context.Background(), params)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.Status)
	assert.Equal(t, "peopledatalabs.com", resp.Data[0].Website)
}
