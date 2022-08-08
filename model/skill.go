package model

import (
	"errors"
)

type SkillBaseParams struct {
	Skill  string           `json:"skill,omitempty" url:"jobskill_title,omitempty"` // Skill that is used as the seed for enrichment
}

type SkillParams struct {
	BaseParams
	SkillBaseParams
	AdditionalParams
}

func (params SkillParams) Validate() error {
	if params.Skill == "" {
		return errors.New("skill: 'Skill' para must not be empty")
	}
	return nil
}

type SkillResponse struct {
	Status int                  `json:"status"`
	Data   SkillResult          `json:"data"`
}

type SkillResult struct {
	CleanedSkill string         `json:"cleaned_skill"`
	SimilarSkills []string   `json:"similar_skills"`
	RelevantJobTitles []string  `json:"relevant_job_titles"`
}