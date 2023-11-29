package cohere

import "encoding/json"

type APIError struct {
	Message string `json:"message"`
}

func (e APIError) Error() string {
	b, err := json.Marshal(e)
	if err != nil {
		return "unknown error"
	}
	return string(b)
}
