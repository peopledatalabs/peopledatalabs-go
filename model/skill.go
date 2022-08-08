package model

import (
	"errors"
)

type SkillBaseParams struct {
	Skill     string   `json:"skill,omitempty" url:"skill,omitempty"` // Skill that is used as the seed for enrichment
	Pretty    bool     `json:"pretty,omitempty" url:"pretty,omitempty"` // Whether the output should have human-readable indentation.
}

type SkillParams struct {
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