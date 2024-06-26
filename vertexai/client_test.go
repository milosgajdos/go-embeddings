package vertexai

import (
	"testing"

	"github.com/milosgajdos/go-embeddings/client"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

const (
	vertexaiToken   = "token"
	vertexaiModel   = "model"
	googleProjectID = "project"
)

type ts struct {
	token string
}

func (t *ts) Token() (*oauth2.Token, error) {
	return &oauth2.Token{AccessToken: t.token}, nil
}

func TestClient(t *testing.T) {
	t.Setenv("VERTEXAI_TOKEN", vertexaiToken)
	t.Setenv("VERTEXAI_MODEL_ID", vertexaiModel)
	t.Setenv("GOOGLE_PROJECT_ID", googleProjectID)

	t.Run("token", func(t *testing.T) {
		c := NewClient()
		assert.Equal(t, c.opts.Token, vertexaiToken)

		testVal := "foo"
		c = NewClient(WithToken(testVal))
		assert.Equal(t, c.opts.Token, testVal)
	})

	t.Run("token source", func(t *testing.T) {
		c := NewClient()
		assert.Nil(t, c.opts.TokenSrc)

		ts := &ts{token: "foo"}
		c = NewClient(WithTokenSrc(ts))
		assert.NotNil(t, c.opts.TokenSrc)
	})

	t.Run("project id", func(t *testing.T) {
		c := NewClient()
		assert.Equal(t, c.opts.ProjectID, googleProjectID)

		testVal := "id"
		c = NewClient(WithProjectID(testVal))
		assert.Equal(t, c.opts.ProjectID, testVal)
	})

	t.Run("model id", func(t *testing.T) {
		c := NewClient()
		assert.Equal(t, c.opts.ModelID, vertexaiModel)

		testVal := "id"
		c = NewClient(WithModelID(testVal))
		assert.Equal(t, c.opts.ModelID, testVal)
	})

	t.Run("BaseURL", func(t *testing.T) {
		c := NewClient()
		assert.Equal(t, c.opts.BaseURL, BaseURL)

		testVal := "http://foo"
		c = NewClient(WithBaseURL(testVal))
		assert.Equal(t, c.opts.BaseURL, testVal)
	})

	t.Run("http client", func(t *testing.T) {
		c := NewClient()
		assert.NotNil(t, c.opts.HTTPClient)

		testVal := client.NewHTTP()
		c = NewClient(WithHTTPClient(testVal))
		assert.NotNil(t, c.opts.HTTPClient)
	})
}
