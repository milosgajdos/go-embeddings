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
	apiKey  string
	baseURL string
	version string
	hc      *http.Client
}

// NewClient creates a new HTTP API client and returns it.
// By default it reads the Cohere API key from COHERE_API_KEY
// env var and uses the default Go http.Client for making API requests.
// You can override the default options via the client methods.
func NewClient() *Client {
	return &Client{
		apiKey:  os.Getenv("COHERE_API_KEY"),
		baseURL: BaseURL,
		version: EmbedAPIVersion,
		hc:      &http.Client{},
	}
}

// NewEmbedder creates a client that implements embeddings.Embedder
func NewEmbedder() embeddings.Embedder[*EmbeddingRequest] {
	return NewClient()
}

// WithAPIKey sets the API key.
func (c *Client) WithAPIKey(apiKey string) *Client {
	c.apiKey = apiKey
	return c
}

// WithBaseURL sets the API base URL.
func (c *Client) WithBaseURL(baseURL string) *Client {
	c.baseURL = baseURL
	return c
}

// WithVersion sets the API version.
func (c *Client) WithVersion(version string) *Client {
	c.version = version
	return c
}

// WithHTTPClient sets the HTTP client.
func (c *Client) WithHTTPClient(httpClient *http.Client) *Client {
	c.hc = httpClient
	return c
}

func (c *Client) doRequest(req *http.Request) (*http.Response, error) {
	resp, err := c.hc.Do(req)
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
