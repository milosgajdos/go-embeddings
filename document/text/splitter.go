package text

import (
	"log"
	"regexp"
)

// Splitter is a character text splitter.
// It splits texts into chunks over
// a separator which is either a string
// or a regular expression.
type Splitter struct {
	Sep          string
	IsSepRegex   bool
	KeepSep      bool
	ChunkSize    int
	ChunkOverlap int
	TrimSpace    bool
	LenFunc      LenFunc
}

func (s *Splitter) mergeChunks(chunks []string, sep string) []string {
	// nolint:prealloc
	var (
		// resulting chunk slice
		resChunks []string
		// buffer of chunks
		chunkBuffer []string
	)

	totalChunksLen := 0
	sepLen := s.LenFunc(sep)

	for _, chunk := range chunks {
		if chunk == "" {
			continue
		}
		splitLen := s.LenFunc(chunk)
		// check if adding this chunk into the buffer execeeds the requested ChunkSize threshold
		// if it does and if the buffer contains any chunks, we'll pop them add them into resulting chunk set.
		if totalChunksLen+splitLen+(sepLen*boolToInt(len(chunkBuffer) > 0)) > s.ChunkSize {
			if totalChunksLen > s.ChunkSize {
				log.Printf("Created a chunk of size %d, which is longer than the requested %d\n", totalChunksLen, s.ChunkSize)
			}
			if len(chunkBuffer) > 0 {
				doc := joinChunks(chunkBuffer, sep, s.TrimSpace)
				if doc != "" {
					resChunks = append(resChunks, doc)
				}
				// Keep on popping chunks from the bffer if:
				// - we have a larger chunk than in the chunk overlap
				// - or if we still have any chunks and the length is long
				for totalChunksLen > s.ChunkOverlap ||
					(totalChunksLen+splitLen+(sepLen*boolToInt(len(chunkBuffer) > 0)) > s.ChunkSize && totalChunksLen > 0) {
					totalChunksLen -= s.LenFunc(chunkBuffer[0]) + sepLen*boolToInt(len(chunkBuffer) > 1)
					chunkBuffer = chunkBuffer[1:]
				}
			}
		}
		chunkBuffer = append(chunkBuffer, chunk)
		// NOTE: we account for any existing chunks in the buffer
		// as joinChunks joins the chunks with the separator string.
		totalChunksLen += splitLen + sepLen*boolToInt(len(chunkBuffer) > 1)
	}

	chunk := joinChunks(chunkBuffer, sep, s.TrimSpace)
	if chunk != "" {
		resChunks = append(resChunks, chunk)
	}

	return resChunks
}

// Split splits text into chunks.
func (s *Splitter) Split(text string) []string {
	sep := s.Sep
	if !s.IsSepRegex {
		sep = regexp.QuoteMeta(s.Sep)
	}
	chunks := splitText(text, sep, s.KeepSep)
	sep = ""
	if !s.KeepSep {
		sep = s.Sep
	}
	return s.mergeChunks(chunks, sep)
}
