package api

import (
	"net/http"
	"testing"

	"github.com/peopledatalabs/peopledatalabs-go/model"

	"github.com/stretchr/testify/assert"
)

func TestCompany_Enrich(t *testing.T) {
	// setup
	server := mockServer()
	defer server.Close()
	company := Company{Client: mockClient(server)}

	// test
	params := model.EnrichCompanyParams{
		CompanyParams: model.CompanyParams{Name: "Google, Inc."},
	}
	resp, err := company.Enrich(params)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.Status)
	assert.Equal(t, 1998, resp.Founded)
}

func TestCompany_Clean(t *testing.T) {
	// setup
	server := mockServer()
	defer server.Close()
	company := Company{Client: mockClient(server)}

	// test
	params := model.CleanCompanyParams{Name: "UCLA"}
	resp, err := company.Clean(params)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, "university of california, los angeles", resp.Name)
}

func TestCompany_Search(t *testing.T) {
	// setup
	server := mockServer()
	defer server.Close()
	company := Company{Client: mockClient(server)}

	// test
	params := model.SearchParams{
		SearchBaseParams: model.SearchBaseParams{
			SQL:  "SELECT * FROM company WHERE name='people data labs'",
			From: 100,
		},
	}
	resp, err := company.Search(params)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.Status)
	assert.Equal(t, "peopledatalabs.com", resp.Data[0].Website)
}
