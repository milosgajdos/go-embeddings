package bedrock

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
)

const (
	// DefaultRegion is default AWS region
	DefaultRegion = "us-east-1"
)

type Client struct {
	opts Options
}

type Options struct {
	Region  string
	ModelID string
	Client  *bedrockruntime.Client
}

// Option is functional option.
type Option func(*Options)

// NewClient creates a new AWS Bedrock HTTP API client and returns it.
// By default it reads the default AWS evnironment variables.
// and constructs the AWS API client.
func NewClient(opts ...Option) *Client {
	options := Options{
		Region:  os.Getenv("AWS_REGION"),
		ModelID: os.Getenv("AWS_BEDROCK_MODEL_ID"),
	}

	for _, apply := range opts {
		apply(&options)
	}

	if options.Region == "" {
		options.Region = DefaultRegion
	}

	if options.Client == nil {
		cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(options.Region))
		if err != nil {
			log.Fatal(err)
		}

		options.Client = bedrockruntime.NewFromConfig(cfg)
	}

	return &Client{
		opts: options,
	}
}

// WithRegion sets AWS region.
func WithRegion(region string) Option {
	return func(o *Options) {
		o.Region = region
	}
}

// WithModelID sets the Tital embedding model ID>
func WithModelID(id string) Option {
	return func(o *Options) {
		o.ModelID = id
	}
}

// WithBedrockClient sets Bedrock API client.
func WithBedrockClient(client *bedrockruntime.Client) Option {
	return func(o *Options) {
		o.Client = client
	}
}
