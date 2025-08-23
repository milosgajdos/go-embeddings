package huggingface

import "encoding/json"

// APIError is error returned by API
type APIError struct {
	Message string `json:"error"`
}

// Error implements errors interface.
func (e APIError) Error() string {
	b, err := json.Marshal(e)
	if err != nil {
		return "unknown error"
	}
	return string(b)
}
