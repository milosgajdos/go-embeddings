package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/milosgajdos/go-embeddings/bedrock"
)

var (
	input string
	model string
)

func init() {
	flag.StringVar(&input, "input", "what is life", "input data")
	flag.StringVar(&model, "model", bedrock.TitanTextV1.String(), "model name")
}

func main() {
	flag.Parse()

	c := bedrock.NewClient(bedrock.WithModelID(model))

	embReq := &bedrock.Request{
		InputText: input,
	}

	embs, err := c.Embed(context.Background(), embReq)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("got %d embeddings", len(embs))
}
