package api

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/peopledatalabs/peopledatalabs-go/v2/model"

	"github.com/stretchr/testify/assert"
)

func TestIP(t *testing.T) {
	// setup
	auto := IP{Client: NewClient(os.Getenv("PDL_API_KEY"), "1.0.0")}

	// test
	params := model.IPParams{
		BaseParams:   model.BaseParams{Pretty: true},
		IPBaseParams: model.IPBaseParams{IP: "72.212.42.169"},
	}
	resp, err := auto.IP(context.Background(), params)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, resp.Status, http.StatusOK)
	assert.Equal(t, resp.Data.IP.Address, "72.212.42.169")
}
