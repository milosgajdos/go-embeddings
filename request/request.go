package request

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/milosgajdos/go-embeddings/client"
)

// NewHTTP creates a new HTTP request from the provided parameters  and returns it.
// If the passed in context is nil, it creates a new background context.
// If the provided body is nil, it gets initialized to bytes.Reader.
// By default the following headers are set:
// * Accept: application/json; charset=utf-8
// If no Content-Type has been set via options it defaults to application/json.
func NewHTTP(ctx context.Context, method, url string, body io.Reader, opts ...Option) (*http.Request, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	if body == nil {
		body = &bytes.Reader{}
	}

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	for _, setOption := range opts {
		setOption(req)
	}

	req.Header.Set("Accept", "application/json; charset=utf-8")
	// if no content-type is specified we default to json
	if ct := req.Header.Get("Content-Type"); len(ct) == 0 {
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	}

	return req, nil
}

// Do sends the HTTP request req using the client and returns the response.
func Do[T error](client *client.HTTP, req *http.Request) (*http.Response, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusBadRequest {
		return resp, nil
	}
	defer resp.Body.Close()

	var apiErr T
	if err := json.NewDecoder(resp.Body).Decode(&apiErr); err != nil {
		return nil, err
	}

	return nil, apiErr
}

// Option is http request functional option.
type Option func(*http.Request)

// WithBearer sets the Authorization header to the provided Bearer token.
func WithBearer(token string) Option {
	return func(req *http.Request) {
		if req.Header == nil {
			req.Header = make(http.Header)
		}
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}
}

// WithSetHeader sets the header key to value val.
func WithSetHeader(key, val string) Option {
	return func(req *http.Request) {
		if req.Header == nil {
			req.Header = make(http.Header)
		}
		req.Header.Set(key, val)
	}
}

// WithAddHeader adds the val to key header.
func WithAddHeader(key, val string) Option {
	return func(req *http.Request) {
		if req.Header == nil {
			req.Header = make(http.Header)
		}
		req.Header.Add(key, val)
	}
}
