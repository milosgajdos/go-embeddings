package openai

import (
	"encoding/json"
	"errors"
)

var (
	// ErrInValidData is returned when the API client fails to decode the returned data.
	ErrInValidData = errors.New("invalid data")
	// ErrUnsupportedEncoding is returned when API client attempts to use unsupported encoding format.
	ErrUnsupportedEncoding = errors.New("unsupported encoding format")
)

// APIError is open AI API error.
type APIError struct {
	Err struct {
		Message string  `json:"message"`
		Type    string  `json:"type"`
		Param   *string `json:"param,omitempty"`
		Code    any     `json:"code,omitempty"`
	} `json:"error"`
}

// Error implements error interface.
func (e APIError) Error() string {
	b, err := json.Marshal(e)
	if err != nil {
		return "unknown error"
	}
	return string(b)
}
