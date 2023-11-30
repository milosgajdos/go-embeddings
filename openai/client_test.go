package openai

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	openaiKey = "t0ps3cr3tk3y"
)

func TestClient(t *testing.T) {
	t.Setenv("OPENAI_API_KEY", openaiKey)

	t.Run("API key", func(t *testing.T) {
		c := NewClient()
		assert.Equal(t, c.apiKey, openaiKey)

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

	t.Run("orgID", func(t *testing.T) {
		c := NewClient()
		assert.Equal(t, c.orgID, "")

		testVal := "orgID"
		c.WithOrgID(testVal)
		assert.Equal(t, c.orgID, testVal)
	})

	t.Run("http client", func(t *testing.T) {
		c := NewClient()
		assert.NotNil(t, c.hc)

		testVal := &http.Client{}
		c.WithHTTPClient(testVal)
		assert.NotNil(t, c.hc)
	})
}
