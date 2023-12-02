package text

import (
	"log"
	"regexp"
	"strings"
)

// Splitter splits text documents.
type Splitter struct {
	chunkSize    int
	chunkOverlap int
	trimSpace    bool
	keepSep      bool
	lenFunc      LenFunc
}

// Config configures the splitter
// NOTE: this is used to prevent situation
// where values in constructors accideentally
// mix the order of parameters of the same type
// leading to unpredicable behaviour.
type Config struct {
	ChunkSize    int
	ChunkOverlap int
	TrimSpace    bool
	KeepSep      bool
	LenFunc      LenFunc
}

// NewSplitterWithConfig creates a new text splitter
// with default options and returns it.
// You can override all config options with appropriate methods.
func NewSplitter() *Splitter {
	return &Splitter{
		chunkSize:    DefaultChunkSize,
		chunkOverlap: DefaultChunkOverlap,
		lenFunc:      DefaultLenFunc,
	}
}

// NewSplitterWithConfig creates a new text splitter and returns it.
func NewSplitterWithConfig(c Config) *Splitter {
	return &Splitter{
		chunkSize:    c.ChunkSize,
		chunkOverlap: c.ChunkOverlap,
		trimSpace:    c.TrimSpace,
		keepSep:      c.KeepSep,
		lenFunc:      c.LenFunc,
	}
}

// WithChunkSize sets chunk size.
func (s *Splitter) WithChunkSize(chunkSize int) *Splitter {
	s.chunkSize = chunkSize
	return s
}

// WithChunkOverlap sets chunk overlap.
func (s *Splitter) WithChunkOverlap(chunkOverlap int) *Splitter {
	s.chunkOverlap = chunkOverlap
	return s
}

// WithTrimSpace sets trim space.
func (s *Splitter) WithTrimSpace(trimSpace bool) *Splitter {
	s.trimSpace = trimSpace
	return s
}

// WithLenFunc sets length func.
func (s *Splitter) WithLenFunc(f LenFunc) *Splitter {
	s.lenFunc = f
	return s
}

// WithKeepSep sets keep separator flag.
func (s *Splitter) WithKeepSep(keepSep bool) *Splitter {
	s.keepSep = keepSep
	return s
}

// join joins the chunks over the given separator into a string
// optionally trimming empty space characters and returns it.
func (s *Splitter) join(chunks []string, sep string) string {
	text := strings.Join(chunks, sep)
	if s.trimSpace {
		text = strings.TrimSpace(text)
	}
	return text
}

// merge merges chunks over the given separator and returns
// the new slice of chunks taking into account chunk overlap.
// It ignores empty string chunks and warns if a chunk is generated
// that exceeds the set chunk size.
func (s *Splitter) merge(chunks []string, sep string) []string {
	// nolint:prealloc
	var (
		// resulting chunk slice
		resChunks []string
		// buffer of chunks
		chunkBuffer []string
	)

	totalChunksLen := 0
	sepLen := s.lenFunc(sep)

	for _, chunk := range chunks {
		if chunk == "" {
			continue
		}
		splitLen := s.lenFunc(chunk)
		// check if adding this chunk into the buffer execeeds the requested chunkSize threshold
		// if it does and if the buffer contains any chunks, we'll pop them add them into resulting chunk set.
		if totalChunksLen+splitLen+(sepLen*boolToInt(len(chunkBuffer) > 0)) > s.chunkSize {
			if totalChunksLen > s.chunkSize {
				log.Printf("Created a chunk of size %d, which is longer than the requested %d\n", totalChunksLen, s.chunkSize)
			}
			if len(chunkBuffer) > 0 {
				doc := s.join(chunkBuffer, sep)
				if doc != "" {
					resChunks = append(resChunks, doc)
				}
				// Keep on popping chunks from the bffer if:
				// - we have a larger chunk than in the chunk overlap
				// - or if we still have any chunks and the length is long
				for totalChunksLen > s.chunkOverlap ||
					(totalChunksLen+splitLen+(sepLen*boolToInt(len(chunkBuffer) > 0)) > s.chunkSize && totalChunksLen > 0) {
					totalChunksLen -= s.lenFunc(chunkBuffer[0]) + sepLen*boolToInt(len(chunkBuffer) > 1)
					chunkBuffer = chunkBuffer[1:]
				}
			}
		}
		chunkBuffer = append(chunkBuffer, chunk)
		// NOTE: we account for any existing chunks in the buffer
		// as s.join joins the chunks with the separator string.
		totalChunksLen += splitLen + sepLen*boolToInt(len(chunkBuffer) > 1)
	}

	chunk := s.join(chunkBuffer, sep)
	if chunk != "" {
		resChunks = append(resChunks, chunk)
	}

	return resChunks
}

// splitText splits the text over a separator optionally keeping
// the separator and returns the the chunks in a slice.
// If the separator is empty string it splits on individual characters.
func (s *Splitter) splitText(text string, sep string) []string {
	if sep != "" {
		if s.keepSep {
			var results []string
			re := regexp.MustCompile("(" + sep + ")")
			splits := re.Split(text, -1)
			for i := 1; i < len(splits); i++ {
				// make sure the separator remains in the result split
				// because Go reasons: https://github.com/golang/go/issues/18868
				results = append(results, sep+splits[i])
			}
			results = append([]string{splits[0]}, results...)
			return results
		}
		re := regexp.MustCompile(sep)
		return re.Split(text, -1)
	}
	// If separator is empty, split into individual characters.
	return strings.Split(text, "")
}

// boolToInt returns 1 if b is true
// and false otherwise.
func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
