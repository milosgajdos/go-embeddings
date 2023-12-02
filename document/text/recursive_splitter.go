package text

import (
	"log"
	"regexp"
)

// RecursiveSplitter is a recursive
// character text splitter.
// It tries to split text recursively  by different
// separators to find one that works.
type RescursiveSplitter struct {
	Seps         []string
	IsSepRegex   bool
	KeepSep      bool
	ChunkSize    int
	ChunkOverlap int
	TrimSpace    bool
	LenFunc      LenFunc
}

func (r *RescursiveSplitter) split(text string, seps []string) []string {
	var (
		resChunks []string
		newSeps   []string
	)

	sep := seps[len(seps)-1]

	for i, s := range seps {
		if !r.IsSepRegex {
			s = regexp.QuoteMeta(s)
		}
		if s == "" {
			sep = s
			break
		}
		if match, _ := regexp.MatchString(s, text); match {
			sep = s
			newSeps = seps[i+1:]
			break
		}
	}

	// TODO should we escape again? Seems weird.
	newSep := sep
	if !r.IsSepRegex {
		newSep = regexp.QuoteMeta(sep)
	}
	chunks := splitText(text, newSep, r.KeepSep)

	var goodChunks []string

	if r.KeepSep {
		newSep = ""
	}

	for _, chunk := range chunks {
		if r.LenFunc(chunk) < r.ChunkSize {
			goodChunks = append(goodChunks, chunk)
			continue
		}

		if len(goodChunks) > 0 {
			mergedText := r.mergeChunks(goodChunks, newSep)
			resChunks = append(resChunks, mergedText...)
			// TODO: reset slice
			goodChunks = nil
		}

		if len(newSeps) == 0 {
			resChunks = append(resChunks, chunk)
			continue
		}

		otherChunks := r.split(chunk, newSeps)
		resChunks = append(resChunks, otherChunks...)
	}

	if len(goodChunks) > 0 {
		mergedText := r.mergeChunks(goodChunks, newSep)
		resChunks = append(resChunks, mergedText...)
	}

	return resChunks
}

// Split splits text into chunks.
func (r *RescursiveSplitter) Split(text string) []string {
	return r.split(text, r.Seps)
}

func (r *RescursiveSplitter) mergeChunks(chunks []string, sep string) []string {
	// nolint:prealloc
	var (
		// resulting chunk slice
		resChunks []string
		// buffer of chunks
		chunkBuffer []string
	)

	totalChunksLen := 0
	sepLen := r.LenFunc(sep)

	for _, chunk := range chunks {
		// TODO: remove all empty char slices actually
		if chunk == "" {
			continue
		}
		splitLen := r.LenFunc(chunk)
		// check if adding this chunk into the buffer execeeds the requested ChunkSize threshold
		// if it does and if the buffer contains any chunks, we'll pop them add them into resulting chunk set.
		if totalChunksLen+splitLen+(sepLen*boolToInt(len(chunkBuffer) > 0)) > r.ChunkSize {
			if totalChunksLen > r.ChunkSize {
				log.Printf("Created a chunk of size %d, which is longer than the requested %d\n", totalChunksLen, r.ChunkSize)
			}
			if len(chunkBuffer) > 0 {
				doc := joinChunks(chunkBuffer, sep, r.TrimSpace)
				if doc != "" {
					resChunks = append(resChunks, doc)
				}
				// Keep on popping chunks from the bffer if:
				// - we have a larger chunk than in the chunk overlap
				// - or if we still have any chunks and the length is long
				for totalChunksLen > r.ChunkOverlap ||
					(totalChunksLen+splitLen+(sepLen*boolToInt(len(chunkBuffer) > 0)) > r.ChunkSize && totalChunksLen > 0) {
					totalChunksLen -= r.LenFunc(chunkBuffer[0]) + sepLen*boolToInt(len(chunkBuffer) > 1)
					chunkBuffer = chunkBuffer[1:]
				}
			}
		}
		chunkBuffer = append(chunkBuffer, chunk)
		// NOTE: we account for separator length in the chunk
		// as joinChunks joins the chunks with the separator string.
		totalChunksLen += splitLen + sepLen*boolToInt(len(chunkBuffer) > 1)
	}

	chunk := joinChunks(chunkBuffer, sep, r.TrimSpace)
	if chunk != "" {
		resChunks = append(resChunks, chunk)
	}

	return resChunks
}
