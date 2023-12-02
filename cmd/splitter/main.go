package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/milosgajdos/go-embeddings/document/text"
)

var (
	input        string
	chunkSize    int
	chunkOverlap int
	trimSpace    bool
	keepSep      bool
)

func init() {
	flag.StringVar(&input, "input", "", "document input")
	flag.IntVar(&chunkSize, "chunk-size", 100, "chunk size")
	flag.IntVar(&chunkOverlap, "chunk-overlap", 10, "chunk overlap")
	flag.BoolVar(&trimSpace, "trim", false, "trim empty space chars from chunks")
	flag.BoolVar(&keepSep, "keep-separator", false, "keep separator in chunks")
}

func main() {
	flag.Parse()

	if input == "" {
		log.Fatal("empty input path")
	}

	content, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}

	s := text.NewSplitter().
		WithChunkSize(chunkSize).
		WithChunkOverlap(chunkOverlap).
		WithTrimSpace(true).
		WithKeepSep(true)

	rs := text.NewRecursiveCharSplitter().
		WithSplitter(s)

	splits := rs.Split(string(content))

	fmt.Println(len(splits))
	for i, s := range splits {
		fmt.Println(i, s)
	}
}
