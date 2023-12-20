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
		assert.Equal(t, c.opts.APIKey, openaiKey)

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

	t.Run("orgID", func(t *testing.T) {
		c := NewClient()
		assert.Equal(t, c.opts.OrgID, "")

		testVal := "orgID"
		c = NewClient(WithOrgID(testVal))
		assert.Equal(t, c.opts.OrgID, testVal)
	})

	t.Run("http client", func(t *testing.T) {
		c := NewClient()
		assert.NotNil(t, c.opts.HTTPClient)

		testVal := &http.Client{}
		c = NewClient(WithHTTPClient(testVal))
		assert.NotNil(t, c.opts.HTTPClient)
	})
}
