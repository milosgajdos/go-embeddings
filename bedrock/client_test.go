package bedrock

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/stretchr/testify/assert"
)

const (
	bedrockModelID = "model"
)

func TestClient(t *testing.T) {
	t.Parallel()

	t.Run("Region", func(t *testing.T) {
		t.Parallel()
		c := NewClient()
		assert.Equal(t, c.opts.Region, DefaultRegion)

		testVal := "us-west-1"
		c = NewClient(WithRegion(testVal))
		assert.Equal(t, c.opts.Region, testVal)
	})

	t.Run("ModelID", func(t *testing.T) {
		t.Parallel()
		c := NewClient()
		assert.Equal(t, c.opts.ModelID, "")

		c = NewClient(WithModelID(bedrockModelID))
		assert.Equal(t, c.opts.ModelID, bedrockModelID)
	})

	t.Run("BedrockClient", func(t *testing.T) {
		t.Parallel()
		c := NewClient()
		assert.NotNil(t, c.opts.Client)
		assert.Equal(t, c.opts.Client.Options().Region, DefaultRegion)

		testVal := "us-west-1"
		cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(testVal))
		assert.NoError(t, err)
		bc := bedrockruntime.NewFromConfig(cfg)

		c = NewClient(WithBedrockClient(bc))
		assert.NotNil(t, c.opts.Client)
		assert.Equal(t, c.opts.Client.Options().Region, testVal)
	})
}
