package vertexai

import (
	"errors"

	"golang.org/x/oauth2"
)

var (
	ErrMissingTokenSource = errors.New("missing access token source")
)

const (
	// NOTE: https://developers.google.com/identity/protocols/oauth2/scopes
	Scopes = "https://www.googleapis.com/auth/cloud-platform"
)

// GetToken returns access token from the given token source.
func GetToken(tokenSrc oauth2.TokenSource) (string, error) {
	if tokenSrc != nil {
		token, err := tokenSrc.Token()
		if err != nil {
			return "", err
		}
		return token.AccessToken, nil
	}

	return "", ErrMissingTokenSource
}
