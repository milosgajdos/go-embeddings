package cohere

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/milosgajdos/go-embeddings"
)

const (
	// BaseURL is Cohere HTTP API base URL.
	BaseURL = "https://api.cohere.ai"
	// EmbedAPIVersion is the latest stable embedding API version.
	EmbedAPIVersion = "v1"
)

// Client is Cohere HTTP API client.
type Client struct {
	opts Options
}

// Options are client options
type Options struct {
	APIKey     string
	BaseURL    string
	Version    string
	HTTPClient *http.Client
}

// Option is functional graph option.
type Option func(*Options)

// NewClient creates a new HTTP API client and returns it.
// By default it reads the Cohere API key from COHERE_API_KEY
// env var and uses the default Go http.Client for making API requests.
// You can override the default options via the client methods.
func NewClient(opts ...Option) *Client {
	options := Options{
		APIKey:     os.Getenv("COHERE_API_KEY"),
		BaseURL:    BaseURL,
		Version:    EmbedAPIVersion,
		HTTPClient: &http.Client{},
	}

	for _, apply := range opts {
		apply(&options)
	}

	return &Client{
		opts: options,
	}
}

// NewEmbedder creates a client that implements embeddings.Embedder
func NewEmbedder(opts ...Option) embeddings.Embedder[*EmbeddingRequest] {
	return NewClient(opts...)
}

// WithAPIKey sets the API key.
func WithAPIKey(apiKey string) Option {
	return func(o *Options) {
		o.APIKey = apiKey
	}
}

// WithBaseURL sets the API base URL.
func WithBaseURL(baseURL string) Option {
	return func(o *Options) {
		o.BaseURL = baseURL
	}
}

// WithVersion sets the API version.
func WithVersion(version string) Option {
	return func(o *Options) {
		o.Version = version
	}
}

// WithHTTPClient sets the HTTP client.
func WithHTTPClient(httpClient *http.Client) Option {
	return func(o *Options) {
		o.HTTPClient = httpClient
	}
}

func (c *Client) doRequest(req *http.Request) (*http.Response, error) {
	resp, err := c.opts.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusBadRequest {
		return resp, nil
	}
	defer resp.Body.Close()

	var apiErr APIError
	if err := json.NewDecoder(resp.Body).Decode(&apiErr); err != nil {
		return nil, err
	}

	return nil, apiErr
}
