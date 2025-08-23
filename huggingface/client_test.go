package huggingface

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	huggingFaceKey = "somekey"
)

func TestClient(t *testing.T) {
	t.Setenv("HUGGINGFACE_API_KEY", huggingFaceKey)

	t.Run("API key", func(t *testing.T) {
		c := NewClient()
		assert.Equal(t, c.apiKey, huggingFaceKey)

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

	t.Run("Model", func(t *testing.T) {
		c := NewClient()
		assert.Equal(t, c.model, "")

		testVal := "foo/bar"
		c.WithModel(testVal)
		assert.Equal(t, c.model, testVal)
	})

	t.Run("http client", func(t *testing.T) {
		c := NewClient()
		assert.NotNil(t, c.hc)

		testVal := &http.Client{}
		c.WithHTTPClient(testVal)
		assert.NotNil(t, c.hc)
	})
}
