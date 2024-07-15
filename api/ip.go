package api

import (
	"context"

	"github.com/peopledatalabs/peopledatalabs-go/v3/model"
)

const (
	ipPath = "/ip/enrich"
)

type IP struct {
	Client
}

type IPFunc func(ctx context.Context, params model.IPParams) (model.IPResponse, error)

// IP allows your users to enrich IP addresses
func (a IP) IP(ctx context.Context, params model.IPParams) (model.IPResponse, error) {
	if err := params.Validate(); err != nil {
		return model.IPResponse{}, err
	}
	var response model.IPResponse
	return response, a.Client.get(ctx, ipPath, params, &response)
}
