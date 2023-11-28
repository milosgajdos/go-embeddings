package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/milosgajdos/go-embedding/openai"
)

var (
	input    string
	model    string
	encoding string
)

func init() {
	flag.StringVar(&input, "input", "", "input data")
	flag.StringVar(&model, "model", string(openai.TextAdaV2), "model name")
	flag.StringVar(&encoding, "encoding", string(openai.EncodingFloat), "encoding format")
}

func main() {
	flag.Parse()

	c, err := openai.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	embReq := &openai.EmbeddingRequest{
		Input:          input,
		Model:          openai.Model(model),
		EncodingFormat: openai.EncodingFormat(encoding),
	}

	embs, err := c.CreateEmbeddings(context.Background(), embReq)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("got %d embeddings", len(embs))
}
