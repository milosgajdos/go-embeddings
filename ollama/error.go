package ollama

import "encoding/json"

// APIError is Ollama API error.
type APIError struct {
	ErrorMessage string `json:"error"`
}

// Error implements errors interface.
func (e APIError) Error() string {
	b, err := json.Marshal(e)
	if err != nil {
		return "unknown error"
	}
	return string(b)
}
