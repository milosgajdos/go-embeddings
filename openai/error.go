package openai

import (
	"encoding/json"
	"errors"
)

var (
	ErrInValidData = errors.New("invalid data")
)

type APIError struct {
	Err struct {
		Message string  `json:"message"`
		Type    string  `json:"type"`
		Param   *string `json:"param,omitempty"`
		Code    any     `json:"code,omitempty"`
	} `json:"error"`
}

func (e APIError) Error() string {
	b, err := json.Marshal(e)
	if err != nil {
		return "unknown error"
	}
	return string(b)
}
