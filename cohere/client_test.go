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
		assert.Equal(t, c.apiKey, cohereAPIKey)

		testVal := "foo"
		c.WithAPIKey(testVal)
		assert.Equal(t, c.apiKey, testVal)
	})

	t.Run("BaseURL", func(t *testing.T) {
		c := NewClient()
		assert.Equal(t, c.baseURL, BaseURL)

		testVal := "http://foo"
		c.WithBaseURL(testVal)
		assert.Equal(t, c.baseURL, testVal)
	})

	t.Run("version", func(t *testing.T) {
		c := NewClient()
		assert.Equal(t, c.version, EmbedAPIVersion)

		testVal := "v3"
		c.WithVersion(testVal)
		assert.Equal(t, c.version, testVal)
	})

	t.Run("http client", func(t *testing.T) {
		c := NewClient()
		assert.NotNil(t, c.hc)

		testVal := &http.Client{}
		c.WithHTTPClient(testVal)
		assert.NotNil(t, c.hc)
	})
}
