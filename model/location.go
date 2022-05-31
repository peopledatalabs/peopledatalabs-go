package model

import "errors"

type LocationParams struct {
	Location string `json:"location,omitempty" url:"location,omitempty"` // The raw location to process
}

type CleanLocationParams struct {
	BaseParams
	LocationParams
	AdditionalParams
}

func (params CleanLocationParams) Validate() error {
	if params.Location == "" {
		return errors.New("location clean: you must provider a location")
	}
	return nil
}

type CleanLocationResponse struct {
	Status int `json:"status"`
	Location
}
