package main

import (
	"flag"
	"fmt"

	"github.com/milosgajdos/go-embeddings/document/text"
)

var content = `What I Worked On

February 2021

Before college the two main things I worked on, outside of school, were writing and programming. I didn't write essays. I wrote what beginning writers were supposed to write then, and probably still are: short stories. My stories were awful. They had hardly any plot, just characters with strong feelings, which I imagined made them deep.

The first programs I tried writing were on the IBM 1401 that our school district used for what was then called "data processing." This was in 9th grade, so I was 13 or 14. The school district's 1401 happened to be in the basement of our junior high school, and my friend Rich Draves and I got permission to use it. It was like a mini Bond villain's lair down there, with all these alien-looking machines — CPU, disk drives, printer, card reader — sitting up on a raised floor under bright fluorescent lights.
`

var (
	chunkSize    int
	chunkOverlap int
	trimSpace    bool
	keepSep      bool
)

func init() {
	flag.IntVar(&chunkSize, "chunk-size", 100, "chunk size")
	flag.IntVar(&chunkOverlap, "chunk-overlap", 10, "chunk overlap")
	flag.BoolVar(&trimSpace, "trim", false, "trim empty space chars from chunks")
	flag.BoolVar(&keepSep, "keep-separator", false, "keep separator in chunks")
}

func main() {
	flag.Parse()

	s := text.NewSplitter().
		WithChunkSize(chunkSize).
		WithChunkOverlap(chunkOverlap).
		WithTrimSpace(true).
		WithKeepSep(true)

	rs := text.NewRecursiveCharSplitter().
		WithSplitter(s)

	splits := rs.Split(content)

	fmt.Println(len(splits))
	for i, s := range splits {
		fmt.Println(i, s)
	}
}
