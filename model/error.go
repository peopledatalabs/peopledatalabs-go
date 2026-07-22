package model

import (
	"encoding/json"
	"fmt"
	"strings"
)

// StringOrStringSlice unmarshals a JSON value that may be either a single
// string or an array of strings, since the PDL API is inconsistent about
// which shape it returns for error.type across endpoints.
type StringOrStringSlice []string

func (s *StringOrStringSlice) UnmarshalJSON(data []byte) error {
	var single string
	if err := json.Unmarshal(data, &single); err == nil {
		*s = StringOrStringSlice{single}
		return nil
	}

	var multiple []string
	if err := json.Unmarshal(data, &multiple); err != nil {
		return err
	}
	*s = StringOrStringSlice(multiple)
	return nil
}

type RestError struct {
	Status  int `json:"status"`
	Details struct {
		Type    StringOrStringSlice `json:"type"`
		Message string              `json:"message"`
	} `json:"error"`
}

func (err RestError) Error() string {
	return fmt.Sprintf("status: %d - error: %s (%s)", err.Status, err.Details.Message, strings.Join(err.Details.Type, ","))
}

type NotFoundError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (err NotFoundError) Error() string {
	return fmt.Sprintf("status: %d - error: %s", err.Status, err.Message)
}
