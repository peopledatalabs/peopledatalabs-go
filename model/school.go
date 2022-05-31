package model

import "errors"

type SchoolParams struct {
	Name    string `json:"name,omitempty" url:"name,omitempty"`       // The name of the school
	Website string `json:"website,omitempty" url:"website,omitempty"` // A website the school uses
	Profile string `json:"profile,omitempty" url:"profile,omitempty"`
}

type CleanSchoolParams struct {
	BaseParams
	SchoolParams
	AdditionalParams
}

func (params CleanSchoolParams) Validate() error {
	if params.Name == "" && params.Website == "" && params.Profile == "" {
		return errors.New("you must input a name OR website OR profile")
	}
	return nil
}

type CleanSchoolResponse struct {
	Status int `json:"status"`
	School
}
