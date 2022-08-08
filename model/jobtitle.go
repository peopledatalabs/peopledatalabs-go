package model

import (
	"errors"
)

type JobTitleBaseParams struct {
	JobTitle  string           `json:"job_title,omitempty" url:"job_title,omitempty"` // JobTitle that is used as the seed for enrichment
}

type JobTitleParams struct {
	BaseParams
	JobTitleBaseParams
	AdditionalParams
}

func (params JobTitleParams) Validate() error {
	if params.JobTitle == "" {
		return errors.New("jobtitle: 'JobTitle' para must not be empty")
	}
	return nil
}

type JobTitleResponse struct {
	Status int                  `json:"status"`
	Data   {}JobTitleResult        `json:"data"`
}

type JobTitleResult struct {
	CleanedJobTitle string      `json:"cleaned_job_title"`
	SimilarJobTitles []string   `json:"similar_job_titles"`
	RelevantSkills []string     `json:"relevant_skills"`
}