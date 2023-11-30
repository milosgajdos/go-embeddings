package vertexai

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

const (
	vertexaiToken   = "token"
	vertexaiMoel    = "model"
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
	t.Setenv("VERTEXAI_MODEL_ID", vertexaiMoel)
	t.Setenv("GOOGLE_PROJECT_ID", googleProjectID)

	t.Run("token", func(t *testing.T) {
		c := NewClient()
		assert.Equal(t, c.token, vertexaiToken)

		testVal := "foo"
		c.WithToken(testVal)
		assert.Equal(t, c.token, testVal)
	})

	t.Run("token source", func(t *testing.T) {
		c := NewClient()
		assert.Nil(t, c.tokenSrc)

		ts := &ts{token: "foo"}
		c.WithTokenSrc(ts)
		assert.NotNil(t, c.tokenSrc)
	})

	t.Run("project id", func(t *testing.T) {
		c := NewClient()
		assert.Equal(t, c.projectID, googleProjectID)

		testVal := "id"
		c.WithProjectID(testVal)
		assert.Equal(t, c.projectID, testVal)
	})

	t.Run("model id", func(t *testing.T) {
		c := NewClient()
		assert.Equal(t, c.modelID, vertexaiMoel)

		testVal := "id"
		c.WithProjectID(testVal)
		assert.Equal(t, c.modelID, vertexaiMoel)
	})

	t.Run("BaseURL", func(t *testing.T) {
		c := NewClient()
		assert.Equal(t, c.baseURL, BaseURL)

		testVal := "http://foo"
		c.WithBaseURL(testVal)
		assert.Equal(t, c.baseURL, testVal)
	})

	t.Run("http client", func(t *testing.T) {
		c := NewClient()
		assert.NotNil(t, c.hc)

		testVal := &http.Client{}
		c.WithHTTPClient(testVal)
		assert.NotNil(t, c.hc)
	})
}
