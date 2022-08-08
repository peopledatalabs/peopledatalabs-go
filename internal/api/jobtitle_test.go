package api

import (
	"net/http"
	"os"
	"testing"

	"github.com/peopledatalabs/peopledatalabs-go/model"

	"github.com/stretchr/testify/assert"
)

func TestJobTitle(t *testing.T) {
	// setup
	auto := JobTitle{Client: NewClient(os.Getenv("PDL_API_KEY"), "1.0.0")}

	// test
	params := model.JobTitleParams{
		JobTitleBaseParams: model.JobTitleBaseParams{JobTitle: "data scientist", Pretty: true},
	}
	resp, err := auto.JobTitle(params)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, resp.Status, http.StatusOK)
	assert.Equal(t, resp.Data.CleanedJobTitle, "data scientist")
}
