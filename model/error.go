package model

import (
	"fmt"
	"strings"
)

type RestError struct {
	Status  int `json:"status"`
	Details struct {
		Type    []string `json:"type"`
		Message string   `json:"message"`
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
