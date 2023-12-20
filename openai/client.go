package openai

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/milosgajdos/go-embeddings"
)

const (
	// BaseURL is OpenAI HTTP API base URL.
	BaseURL = "https://api.openai.com"
	// EmbedAPIVersion is the latest stable embeddings API version.
	EmbedAPIVersion = "v1"
	// OrgHeader is an Organization header
	OrgHeader = "OpenAI-Organization"
)

// Client is an OpenAI HTTP API client.
type Client struct {
	opts Options
}

type Options struct {
	APIKey     string
	BaseURL    string
	Version    string
	OrgID      string
	HTTPClient *http.Client
}

// Option is functional graph option.
type Option func(*Options)

// NewClient creates a new HTTP API client and returns it.
// By default it reads the OpenAI API key from OPENAI_API_KEY
// env var and uses the default Go http.Client for making API requests.
// You can override the default options via the client methods.
func NewClient(opts ...Option) *Client {
	options := Options{
		APIKey:     os.Getenv("OPENAI_API_KEY"),
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

// WithOrgID sets the organization ID.
func WithOrgID(orgID string) Option {
	return func(o *Options) {
		o.OrgID = orgID
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
