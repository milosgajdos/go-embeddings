package cohere

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	cohereAPIKey = "somekey"
)

func TestClient(t *testing.T) {
	t.Setenv("COHERE_API_KEY", cohereAPIKey)

	t.Run("API key", func(t *testing.T) {
		c := NewClient()
		assert.Equal(t, c.opts.APIKey, cohereAPIKey)

		testVal := "foo"
		c = NewClient(WithAPIKey(testVal))
		assert.Equal(t, c.opts.APIKey, testVal)
	})

	t.Run("BaseURL", func(t *testing.T) {
		c := NewClient()
		assert.Equal(t, c.opts.BaseURL, BaseURL)

		testVal := "http://foo"
		c = NewClient(WithBaseURL(testVal))
		assert.Equal(t, c.opts.BaseURL, testVal)
	})

	t.Run("version", func(t *testing.T) {
		c := NewClient()
		assert.Equal(t, c.opts.Version, EmbedAPIVersion)

		testVal := "v3"
		c = NewClient(WithVersion(testVal))
		assert.Equal(t, c.opts.Version, testVal)
	})

	t.Run("http client", func(t *testing.T) {
		c := NewClient()
		assert.NotNil(t, c.opts.HTTPClient)

		testVal := &http.Client{}
		c = NewClient(WithHTTPClient(testVal))
		assert.NotNil(t, c.opts.HTTPClient)
	})
}
