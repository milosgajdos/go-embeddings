package cohere

import "encoding/json"

// APIError is Cohere API error.
type APIError struct {
	Message string `json:"message"`
}

// Error implements errors interface.
func (e APIError) Error() string {
	b, err := json.Marshal(e)
	if err != nil {
		return "unknown error"
	}
	return string(b)
}
