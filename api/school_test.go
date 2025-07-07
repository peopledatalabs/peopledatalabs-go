package api

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/peopledatalabs/peopledatalabs-go/v6/model"

	"github.com/stretchr/testify/assert"
)

func TestSchool_Clean(t *testing.T) {
	// setup
	school := School{Client: NewClient(os.Getenv("PDL_API_KEY"), "1.0.0")}

	// test
	params := model.CleanSchoolParams{
		SchoolParams: model.SchoolParams{Profile: "linkedin.com/school/ucla"},
	}
	resp, err := school.Clean(context.Background(), params)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.Status)
	assert.Equal(t, "ucla.edu", resp.Website)
}
