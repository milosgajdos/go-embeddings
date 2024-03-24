package voyage

import (
	"testing"

	"github.com/milosgajdos/go-embeddings/client"
	"github.com/stretchr/testify/assert"
)

const (
	voyageAPIKey = "somekey"
)

func TestClient(t *testing.T) {
	t.Setenv("VOYAGE_API_KEY", voyageAPIKey)

	t.Run("API key", func(t *testing.T) {
		c := NewClient()
		assert.Equal(t, c.opts.APIKey, voyageAPIKey)

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

		testVal := client.NewHTTP()
		c = NewClient(WithHTTPClient(testVal))
		assert.NotNil(t, c.opts.HTTPClient)
	})
}
