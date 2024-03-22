package client

import (
	"context"
	"net/http"
)

// HTTP is a HTTP client.
type HTTP struct {
	c  *http.Client
	rl Limiter
}

// Options are client options
type Options struct {
	HTTPClient *http.Client
	Limiter    Limiter
}

// Option is functional graph option.
type Option func(*Options)

// Limiter is used to apply rate limits.
// NOTE: you can use off the shelf limiter from
// https://pkg.go.dev/golang.org/x/time/rate#Limiter
type Limiter interface {
	// Wait must block until limiter
	// permits another request to proceed.
	Wait(context.Context) error
}

// NewHTTP creates a new HTTP client and returns it.
func NewHTTP(opts ...Option) *HTTP {
	options := Options{
		HTTPClient: &http.Client{},
	}
	for _, apply := range opts {
		apply(&options)
	}

	return &HTTP{
		c:  options.HTTPClient,
		rl: options.Limiter,
	}
}

// Do dispatches the HTTP request to the network
func (c *HTTP) Do(req *http.Request) (*http.Response, error) {
	err := c.rl.Wait(req.Context()) // This is a blocking call. Honors the rate limit
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// WithHTTPClient sets the HTTP client.
func WithHTTPClient(c *http.Client) Option {
	return func(o *Options) {
		o.HTTPClient = c
	}
}

// WithLimiter sets the http rate limiter.
func WithLimiter(l Limiter) Option {
	return func(o *Options) {
		o.Limiter = l
	}
}
