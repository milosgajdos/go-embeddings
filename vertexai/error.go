package vertexai

import "encoding/json"

// APIError is API error.
type APIError struct {
	RespError struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Status  string `json:"status"`
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
