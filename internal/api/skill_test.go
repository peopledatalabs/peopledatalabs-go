package api

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/peopledatalabs/peopledatalabs-go/model"

	"github.com/stretchr/testify/assert"
)

func TestSkill(t *testing.T) {
	// setup
	auto := Skill{Client: NewClient(os.Getenv("PDL_API_KEY"), "1.0.0")}

	// test
	params := model.SkillParams{
		BaseParams:             model.BaseParams{Pretty: true},
		SkillBaseParams: model.SkillBaseParams{Skill: "python"},
	}
	resp, err := auto.Skill(context.Background(), params)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, resp.Status, http.StatusOK)
	assert.Equal(t, resp.Data.CleanedSkill, "python")
}
