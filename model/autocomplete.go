package model

import (
	"errors"
)

type AutocompleteBaseParams struct {
	Field AutocompleteType `json:"field" url:"field"`                   // Field that autocomplete will be calculated for
	Text  string           `json:"text,omitempty" url:"text,omitempty"` // Text that is used as the seed for autocompletion
}

type AutocompleteParams struct {
	BaseParams
	AutocompleteBaseParams
	AdditionalParams
}

func (params AutocompleteParams) Validate() error {
	if params.Field == "" {
		return errors.New("autocomplete: 'Field' para must not be empty")
	}
	return nil
}

type AutocompleteResponse struct {
	Status int                  `json:"status"`
	Data   []AutocompleteResult `json:"data"`
	Fields []string             `json:"fields"`
}

type AutocompleteResult struct {
	Name  string            `json:"name"`  // The plain text name of this Autocomplete API suggestion. The prefix of this field will match the value of the text input parameter.
	Count int               `json:"count"` // The number of records in our Person Dataset for this Autocomplete API suggestion. This field is used for sorting elements in the data array.
	Meta  map[string]string `json:"meta"`  // A set of additional fields returned for each result in the data array. The metadata fields depend on the field input parameter
}

type AutocompleteType string

const (
	AutocompleteTypeCompany  AutocompleteType = "company"  // Company names
	AutocompleteTypeCountry  AutocompleteType = "country"  // Country names
	AutocompleteTypeIndustry AutocompleteType = "industry" // Industries
	AutocompleteTypeLocation AutocompleteType = "location" // Location names
	AutocompleteTypeMajor    AutocompleteType = "major"    // Educational majors (field of study)
	AutocompleteTypeRegion   AutocompleteType = "region"   // Region name (e.g. states for the US)
	AutocompleteTypeRole     AutocompleteType = "role"     // Job roles
	AutocompleteTypeSubRole  AutocompleteType = "sub_role" // Job sub roles
	AutocompleteTypeSchool   AutocompleteType = "school"   // School names
	AutocompleteTypeSkill    AutocompleteType = "skill"    // Skills
	AutocompleteTypeTitle    AutocompleteType = "title"    // Job titles
	AutocompleteTypeWebsite  AutocompleteType = "website"  // Websites
)
