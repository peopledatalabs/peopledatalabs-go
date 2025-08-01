package api

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/peopledatalabs/peopledatalabs-go/v6/model"

	"github.com/stretchr/testify/assert"
)

func TestLocation_Clean(t *testing.T) {
	// setup
	location := Location{Client: NewClient(os.Getenv("PDL_API_KEY"), "1.0.0")}

	// test
	params := model.CleanLocationParams{
		LocationParams: model.LocationParams{Location: "portland"},
	}
	resp, err := location.Clean(context.Background(), params)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.Status)
	assert.Equal(t, "oregon", resp.Region)
}
