package huggingface

import (
	"encoding/json"
	"net/http"
	"os"
)

const (
	// BaseURL is Cohere HTTP API base URL.
	BaseURL = "https://api-inference.huggingface.co/models"
)

// Client is Cohere HTTP API client.
type Client struct {
	apiKey  string
	baseURL string
	model   string
	hc      *http.Client
}

// NewClient creates a new HTTP API client and returns it.
// By default it reads the Cohere API key from HUGGINGFACE_API_KEY
// env var and uses the default Go http.Client for making API requests.
// You can override the default options via the client methods.
func NewClient() *Client {
	return &Client{
		apiKey:  os.Getenv("HUGGINGFACE_API_KEY"),
		baseURL: BaseURL,
		hc:      &http.Client{},
	}
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

// WithModel sets the model name
func (c *Client) WithModel(model string) *Client {
	c.model = model
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
