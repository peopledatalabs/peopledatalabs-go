package api

import (
	"context"
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
		CompanyParams: model.CompanyParams{Name: "Google, Inc."},
	}
	resp, err := company.Enrich(context.Background(), params)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.Status)
	assert.Equal(t, 1998, resp.Founded)
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
