package api

import (
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
		SkillBaseParams: model.SkillBaseParams{Skill: "c++", Pretty: true},
	}
	resp, err := auto.Skill()

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, resp.Status, http.StatusOK)
	assert.Equal(t, resp.Data.CleanedSkill, "c")
}
