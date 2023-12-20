package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/milosgajdos/go-embeddings/openai"
)

var (
	input    string
	model    string
	encoding string
)

func init() {
	flag.StringVar(&input, "input", "what is life", "input data")
	flag.StringVar(&model, "model", openai.TextAdaV2.String(), "model name")
	flag.StringVar(&encoding, "encoding", openai.EncodingFloat.String(), "encoding format")
}

func main() {
	flag.Parse()

	c := openai.NewClient()

	embReq := &openai.EmbeddingRequest{
		Input:          input,
		Model:          openai.Model(model),
		EncodingFormat: openai.EncodingFormat(encoding),
	}

	embResp, err := c.Embeddings(context.Background(), embReq)
	if err != nil {
		log.Fatal(err)
	}

	embs, err := embResp.ToEmbeddings()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("got %d embeddings", len(embs))
}
