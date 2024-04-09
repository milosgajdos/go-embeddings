package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/milosgajdos/go-embeddings/ollama"
)

var (
	prompt string
	model  string
)

func init() {
	flag.StringVar(&prompt, "prompt", "what is life", "input prompt")
	flag.StringVar(&model, "model", "", "model name")
}

func main() {
	flag.Parse()

	if model == "" {
		log.Fatal("missing ollama model")
	}

	c := ollama.NewClient()

	embReq := &ollama.EmbeddingRequest{
		Prompt: prompt,
		Model:  model,
	}

	embs, err := c.Embed(context.Background(), embReq)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("got %d embeddings", len(embs))
}
