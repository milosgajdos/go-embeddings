package voyage

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

// APIError is Cohere API error.
type APIError struct {
	Detail string `json:"detail"`
}

// Error implements error interface.
func (e APIError) Error() string {
	b, err := json.Marshal(e)
	if err != nil {
		return "unknown error"
	}
	return string(b)
}
