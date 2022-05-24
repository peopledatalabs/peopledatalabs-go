package model

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

type CleanSchoolResponse struct {
	Status int `json:"status"`
	School
}
