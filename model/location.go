package model

type LocationParams struct {
	Location string `json:"location,omitempty" url:"location,omitempty"` // The raw location to process
}

type CleanLocationParams struct {
	BaseParams
	LocationParams
	AdditionalParams
}

type CleanLocationResponse struct {
	Status int `json:"status"`
	Location
}
