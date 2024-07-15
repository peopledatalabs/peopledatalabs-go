package api

import (
	"context"

	"github.com/peopledatalabs/peopledatalabs-go/v3/model"
)

const (
	skillPath = "/skill/enrich"
)

type Skill struct {
	Client
}

type SkillFunc func(ctx context.Context, params model.SkillParams) (model.SkillResponse, error)

// Skill allows your users to get information for Skill values
func (a Skill) Skill(ctx context.Context, params model.SkillParams) (model.SkillResponse, error) {
	if err := params.Validate(); err != nil {
		return model.SkillResponse{}, err
	}
	var response model.SkillResponse
	return response, a.Client.get(ctx, skillPath, params, &response)
}
