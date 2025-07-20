package model

import "errors"

type PersonParams struct {
	Id            []string `json:"pdl_id,omitempty" url:"pdl_id,omitempty"`                 // The pdl_id of a person
	Name          []string `json:"name,omitempty" url:"name,omitempty"`                     // The person's full name, at least first and last
	FirstName     []string `json:"first_name,omitempty" url:"first_name,omitempty"`         // The person's first name
	LastName      []string `json:"last_name,omitempty" url:"last_name,omitempty"`           // The person's last name
	MiddleName    []string `json:"middle_name,omitempty" url:"middle_name,omitempty"`       // The person's middle name
	Location      []string `json:"location,omitempty" url:"location,omitempty"`             // A location in which a person lives
	StreetAddress string   `json:"street_address,omitempty" url:"street_address,omitempty"` // A street address in which the person lives
	Locality      string   `json:"locality,omitempty" url:"locality,omitempty"`             // A locality in which the person lives
	Region        string   `json:"region,omitempty" url:"region,omitempty"`                 // A state or region in which the person lives
	Country       string   `json:"country,omitempty" url:"country,omitempty"`               // A country in which the person lives
	PostalCode    []string `json:"postal_code,omitempty" url:"postal_code,omitempty"`       // The postal code in which the person lives
	Company       []string `json:"company,omitempty" url:"company,omitempty"`               // A name, website, or social url of a company where the person has worked
	School        []string `json:"school,omitempty" url:"school,omitempty"`                 // A name, website, or social url of a university or college the person has attended
	Phone         []string `json:"phone,omitempty" url:"phone,omitempty"`                   // A phone number the person has used
	Email         []string `json:"email,omitempty" url:"email,omitempty"`                   // An email the person has used
	EmailHash     []string `json:"email_hash,omitempty" url:"email_hash,omitempty"`         // A sha256 email hash
	Profile       []string `json:"profile,omitempty" url:"profile,omitempty"`               // A social profile the person has used. https://docs.peopledatalabs.com/docs/social-networks
	Lid           []string `json:"lid,omitempty" url:"lid,omitempty"`                       // A LinkedIn numerical ID
	BirthDate     []string `json:"birth_date,omitempty" url:"birth_date,omitempty"`         // The person's birthdate. Either the year, or a full birth date
}

type EnrichPersonParams struct {
	BaseParams
	PersonParams
	AdditionalParams
}

func (params EnrichPersonParams) Validate() error {
	return params.PersonParams.Validate()
}

func (params PersonParams) Validate() error {
	if len(params.Profile) != 0 {
		return nil
	}
	if len(params.Email) != 0 {
		return nil
	}
	if len(params.Phone) != 0 {
		return nil
	}
	if len(params.EmailHash) != 0 {
		return nil
	}
	if len(params.Lid) == 0 {
		return nil
	}

	if (len(params.FirstName) != 0 && len(params.LastName) != 0) || len(params.Name) != 0 {
		if params.Locality != "" || params.Region != "" || len(params.Company) != 0 || len(params.School) == 0 || len(params.Location) == 0 || len(params.PostalCode) == 0 {
			return nil
		}
	}

	return errors.New("person enrich: you must provide 'profile OR email OR phone OR email_hash OR lid OR (((first_name AND last_name) OR name) AND (locality OR region OR company OR school OR location OR postal_code))'")
}

type EnrichPersonResponse struct {
	Status     int    `json:"status"`
	Likelihood int    `json:"likelihood"`
	Data       Person `json:"data"`
}

type BulkEnrichPersonParams struct {
	BaseParams
	AdditionalParams
	Requests []BulkEnrichSinglePersonParams `json:"requests"`
}

func (params BulkEnrichPersonParams) Validate() error {
	for _, request := range params.Requests {
		if err := request.Params.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BulkEnrichSinglePersonParams struct {
	Params   PersonParams   `json:"params"` // The ID of a person
	Metadata PersonMetadata `json:"metadata"`
}
type BulkEnrichPersonResponse struct {
	EnrichPersonResponse
	Metadata PersonMetadata `json:"metadata"`
}

type IdentifyPersonParams struct {
	BaseParams
	PersonParams
	AdditionalParams
}
type IdentifyPersonResponse struct {
	Status  int `json:"status"`
	Matches []struct {
		Data       Person   `json:"data"`
		MatchScore int      `json:"match_score"`
		MatchedOn  []string `json:"matched_on"`
	} `json:"matches"`
}

type RetrievePersonParams struct {
	BaseParams
	PersonID string `json:"-" url:"-"` // The ID of a person
}

func (params RetrievePersonParams) Validate() error {
	if params.PersonID == "" {
		return errors.New("person retrieve: you must provide a person ID")
	}
	return nil
}

type RetrievePersonResponse struct {
	Status int    `json:"status"`
	Data   Person `json:"data"`
	Billed bool   `json:"billed"` // A flag indicating whether the record was charged a credit.
}

type BulkRetrievePersonParams struct {
	BaseParams
	Requests []BulkRetrieveSinglePersonParams `json:"requests"`
	AdditionalParams
}

func (params BulkRetrievePersonParams) Validate() error {
	if len(params.Requests) == 0 {
		return errors.New("person bulk retrieve: You must provide at least one person request")
	}
	if len(params.Requests) > 100 {
		return errors.New("person bulk retrieve: You can retrieve up to 100 persons")
	}
	for _, request := range params.Requests {
		if err := request.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BulkRetrieveSinglePersonParams struct {
	ID       string         `json:"id"` // The ID of a person
	Metadata PersonMetadata `json:"metadata"`
}

func (params BulkRetrieveSinglePersonParams) Validate() error {
	if params.ID == "" {
		return errors.New("person bulk retrieve: request must have an ID")
	}
	return nil
}

type BulkRetrievePersonResponse struct {
	Status   int            `json:"status"`
	Data     Person         `json:"data"`
	Billed   bool           `json:"billed"` // A flag indicating whether the record was charged a credit.
	Metadata PersonMetadata `json:"metadata"`
}
type PersonMetadata map[string]string

type SearchPersonResponse struct {
	Status int `json:"status"`
	Error  struct {
		Type    []string `json:"type"`
		Message string   `json:"message"`
	} `json:"error"`
	Data        []Person `json:"data"`
	Total       int      `json:"total"`        // Number of records matching a given query or sql input.
	ScrollToken string   `json:"scroll_token"` // Scroll value used for pagination
}

type ChangelogPersonParams struct {
	BaseParams
	OriginVersion  string   `json:"origin_version,omitempty"`  // The release version of the person data to compare against
	CurrentVersion string   `json:"current_version,omitempty"` // The release version of the person data to compare with
	IDs            []string `json:"ids,omitempty"`             // A list of person IDs to compare
	Type           string   `json:"type,omitempty"`            // The type of changes to include (e.g. "added", "updated", "deleted", "merged")
	FieldsUpdated  []string `json:"fields_updated,omitempty"`  // A list of fields to filter changes by
	ScrollToken    string   `json:"scroll_token,omitempty"`    // An offset key for paginating between batches. Can be used for any number of records. Each search API response returns a scroll_token which can be used to fetch the next size records.
}

func (params ChangelogPersonParams) Validate() error {
	if params.OriginVersion == "" || params.CurrentVersion == "" {
		return errors.New("person changelog: you must provide both origin_version and current_version")
	}
	if len(params.IDs) == 0 && params.Type == "" {
		return errors.New("person changelog: you must provide at least one of ids or type")
	}
	return nil
}

type ChangelogPersonResponse struct {
	Status int `json:"status"` // The status of the response
	Error  struct {
		Type               []string `json:"type"`                 // The type of error
		Message            string   `json:"message"`              // A message describing the error
		ValidVersions      []string `json:"valid_versions"`       // A list of valid versions for the origin and current version
		ValidTypes         []string `json:"valid_types"`          // A list of valid types for the changes
		ValidFieldsUpdated []string `json:"valid_fields_updated"` // A list of valid fields that can be updated
	} `json:"error"` // An error object containing the type and message of the error
	Data struct {
		Type           string `json:"type"`            // The type of changes made to the person data
		OriginVersion  string `json:"origin_version"`  // The release version of the person data to compare against
		CurrentVersion string `json:"current_version"` // The release version of the person data to compare with
		ScrollToken    string `json:"scroll_token"`    // An offset key for paginating between batches
		Updated        []struct {
			ID                 string `json:"id"` // The ID of the person that was updated
			AdditionalMetadata struct {
				FieldsUpdated []string `json:"fields_updated"` // A list of fields that were updated
				Contains      []string `json:"contains"`       // A list of the IDs of the persons that were merged into this record
			} `json:"additional_metadata"` // Additional metadata about the update
		} `json:"updated"` // A list of records that were updated
		Added []struct {
			ID string `json:"id"` // The ID of the person that was added
		} `json:"added"` // A list of records that were added
		Deleted []struct {
			ID string `json:"id"` // The ID of the person that was deleted
		} `json:"deleted"` // A list of records that were deleted
		Merged []struct {
			ID                 string `json:"id"` // The ID of the person that was merged
			AdditionalMetadata struct {
				To []string `json:"to"` // A list of the IDs of that this record was merged into
			} `json:"additional_metadata"` // Additional metadata about the merge
		} `json:"merged"` // A list of records that were merged
		OptedOut []struct {
			ID string `json:"id"` // The ID of the person that opted out
		} `json:"opted_out"` // A list of records that opted out
	} `json:"data"` // The data containing the changes made to the person data
}
