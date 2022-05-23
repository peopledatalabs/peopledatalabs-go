package model

type errorResponse struct {
	Status int `json:"status"`
	Error  struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	} `json:"error"`
}
