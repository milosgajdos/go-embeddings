package text

import (
	"errors"
	"log"
	"regexp"
	"regexp/syntax"
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
// NOTE: this is used to prevent situations
// where values in constructors accidentally
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
		trimSpace:    true,
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

// merge merges splits over the given separator and returns
// the new slice of chunks taking into account chunk overlap.
// It ignores empty string chunks and warns if a chunk is generated
// that exceeds the set chunk size.
func (s *Splitter) merge(splits []string, sep Sep) []string {
	// nolint:prealloc
	var (
		// resulting chunk slice
		chunks []string
		// buffer of chunks
		chunkBuffer []string
	)

	totalChunksLen := 0
	sepLen := s.lenFunc(sep.Value)

	for _, chunk := range splits {
		if chunk == "" {
			continue
		}
		splitLen := s.lenFunc(chunk)
		// check if adding this chunk into the buffer execeeds the requested chunkSize threshold
		// if it does and if the buffer contains any chunks, we'll pop them add them into resulting chunk set.
		if totalChunksLen+splitLen+(sepLen*boolToInt(len(chunkBuffer) > 0)) > s.chunkSize {
			if totalChunksLen > s.chunkSize {
				log.Printf("created chunk is longer (%d) than requested: %d\n", totalChunksLen, s.chunkSize)
			}
			if len(chunkBuffer) > 0 {
				doc := s.join(chunkBuffer, sep.Value)
				if doc != "" {
					chunks = append(chunks, doc)
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

	chunk := s.join(chunkBuffer, sep.Value)
	if chunk != "" {
		chunks = append(chunks, chunk)
	}

	return chunks
}

// Split splits the text over a separator optionally keeping
// the separator and returns the the chunks in a slice.
// If the separator is empty string it splits on individual characters.
// TODO: rename this to Split
func (s *Splitter) Split(text string, sep Sep) []string {
	if sep.Value != "" {
		if s.keepSep {
			// NOTE: we must do this to unescape
			// the escaped separator so we keep the raw separator.
			sepVal, err := unquoteMeta(sep.Value)
			if err != nil {
				panic(err)
			}

			var results []string
			splits := regexp.MustCompile("("+sep.Value+")").Split(text, -1)
			// NOTE: we start iterating from 1, not 0!
			for i := 1; i < len(splits); i++ {
				// make sure the separator remains in the result split
				// because Go reasons: https://github.com/golang/go/issues/18868
				results = append(results, sepVal+splits[i])
			}
			results = append([]string{splits[0]}, results...)
			return filterEmptyStrings(results)
		}
		return filterEmptyStrings(regexp.MustCompile(sep.Value).Split(text, -1))
	}
	// If separator is empty, split into individual characters.
	return filterEmptyStrings(strings.Split(text, ""))
}

// boolToInt returns 1 if b is true
// and false otherwise.
func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

// filterEmptyStrings removes empty strings from a slice of strings.
func filterEmptyStrings(slice []string) []string {
	count := 0
	for _, s := range slice {
		if s != "" {
			count++
		}
	}

	result := make([]string, 0, count)

	for _, s := range slice {
		if s != "" {
			result = append(result, s)
		}
	}

	return result
}

// unQuote regexp string.
func unquoteMeta(s string) (string, error) {
	r, err := syntax.Parse(s, 0)
	if err != nil {
		return "", err
	}
	if r.Op != syntax.OpLiteral {
		return "", errors.New("not a quoted meta")
	}
	return string(r.Rune), nil
}
