package api

import (
	"context"
	"net/http"
	"testing"

	"github.com/peopledatalabs/peopledatalabs-go/model"

	"github.com/stretchr/testify/assert"
)

func TestLocation_Clean(t *testing.T) {
	// setup
	server := mockServer()
	defer server.Close()
	location := Location{Client: mockClient(server)}

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
