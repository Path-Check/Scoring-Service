package handler

import (
	"encoding/json"
	"errors"

	"github.com/Path-Check/Scoring-Service/model"
)

var (
	ErrRequestBodyEmpty = errors.New("HTTP request body is empty")
	ErrUnmarshalJSON    = errors.New("Failed to unmarshal into JSON.")
	ErrMarshalJSON      = errors.New("Failed to marshal into JSON.")
)

// This is a handler that can be called by any cloud product specific handler
// once cloud specific modifications have been done.
// Returns: HTTP status code, response body, and error.
func GenericScoringHandler(requestBody string) (int, string, error) {
	// If the HTTP request body is empty, throw an error.
	if len(requestBody) == 0 {
		return 400, "", ErrRequestBodyEmpty
	}

	enReq := model.ExposureNotificationRequest{}
	err := json.Unmarshal([]byte(requestBody), &enReq)
	if err != nil {
		return 400, "", ErrUnmarshalJSON
	}

	enRes, err := model.ScoreV1(&enReq)
	if err != nil {
		return 400, "", err
	}

	response, err := json.Marshal(&enRes)
	if err != nil {
		return 400, "", ErrMarshalJSON
	}

	return 200, string(response), nil
}
