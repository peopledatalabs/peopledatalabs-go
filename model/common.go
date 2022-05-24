package model

type BaseParams struct {
	Pretty bool `json:"pretty,omitempty" url:"pretty,omitempty"` // Whether the output should have human-readable indentation.
	Size   int  `json:"size,omitempty" url:"size,omitempty"`     // The number of matched records to return for this query if they exist*. Must be between 1 and 1000 (inclusive) // TODO: restrict size 1-100
}

type AdditionalParams struct {
	MinLikelihood    int32  `json:"min_likelihood,omitempty" url:"min_likelihood,omitempty"` // The minimum likelihood score a response must possess in order to return a 200
	Required         string `json:"required,omitempty" url:"required,omitempty"`             // Parameter specifying the fields and data points a response must have to return a 200
	TitleCase        bool   `json:"titlecase,omitempty" url:"titlecase,omitempty"`           // Setting titlecase to true will titlecase the person data in 200 responses.
	DataInclude      string `json:"data_include,omitempty" url:"data_include,omitempty"`     // A comma-separated string of fields that you would like the response to include. eg. "names.clean,emails.address". Begin the string with a - if you would instead like to exclude the specified fields. If you would like to exclude all data from being returned, use dataInclude="".
	IncludeIfMatched bool   `json:"include_if_matched" url:"include_if_matched,omitempty"`   // If set to true, includes a top-level (alongside "data", "status", etc) field "matched" which includes a value for each queried field parameter that was "matched-on" during our internal query.
}

type SearchBaseParams struct {
	Query       interface{} `json:"query,omitempty"`                     // An Elasticsearch (v7.7) query. Underlying Elasticsearch mapping reference (https://docs.peopledatalabs.com/docs/search-api#full-field-mapping)
	SQL         string      `json:"sql,omitempty"`                       // A SQL query of the format: SELECT * FROM person WHERE XXX, where XXX is a standard SQL boolean query involving PDL fields. Any use of column selections or the LIMIT keyword will be ignored.
	From        int32       `json:"from,omitempty"`                      // [LEGACY] An offset value for pagination. Can be a number between 0 and 9999. Pagination can be executed up to a maximum of 10,000 records per query. FROM CAN NOT BE USED WITH SCROLL_TOKEN IN THE SAME REQUEST
	ScrollToken string      `json:"scroll_token,omitempty"`              // An offset key for paginating between batches. Can be used for any number of records. Each search API response returns a scroll_token which can be used to fetch the next size records.
	Dataset     string      `json:"dataset,omitempty"`                   // The dataset category to return records from. Can be multiple comma seperated categories or all
	TitleCase   bool        `json:"titlecase" url:"titlecase,omitempty"` // Setting titlecase to true will titlecase the person data in 200 responses.
}

type SearchParams struct {
	BaseParams
	SearchBaseParams
	AdditionalParams
	// TODO: Add validations of min params
}
