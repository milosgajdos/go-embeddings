package text

import (
	"regexp"
)

// RecursiveCharSplitter is a recursive
// character text splitter.
// It tries to split text recursively  by different
// separators to find one that works.
type RecursiveCharSplitter struct {
	*Splitter
	seps []Sep
}

// NewSplitter creates a new splitter and returns it.
func NewRecursiveCharSplitter() *RecursiveCharSplitter {
	return &RecursiveCharSplitter{
		Splitter: NewSplitter(),
		seps:     DefaultSeparators,
	}
}

// WithSplitter sets the splitter.
func (r *RecursiveCharSplitter) WithSplitter(splitter *Splitter) *RecursiveCharSplitter {
	r.Splitter = splitter
	return r
}

// WithSeps sets separators.
func (r *RecursiveCharSplitter) WithSeps(seps []Sep) *RecursiveCharSplitter {
	r.seps = seps
	return r
}

func (r *RecursiveCharSplitter) split(text string, seps []Sep) []string {
	var (
		resChunks []string
		newSeps   []Sep
	)

	sep := seps[len(seps)-1]

	for i, s := range seps {
		if !s.IsRegexp {
			s.Value = regexp.QuoteMeta(s.Value)
		}
		if s.Value == "" {
			sep = s
			break
		}
		if match, _ := regexp.MatchString(s.Value, text); match {
			sep = s
			newSeps = seps[i+1:]
			break
		}
	}

	// TODO should we escape again? Seems weird.
	newSep := Sep{Value: sep.Value, IsRegexp: sep.IsRegexp}
	if !sep.IsRegexp {
		newSep.Value = regexp.QuoteMeta(sep.Value)
	}
	chunks := r.Splitter.Split(text, newSep)

	var goodChunks []string

	if r.keepSep {
		newSep.Value = ""
	}

	for _, chunk := range chunks {
		if r.lenFunc(chunk) < r.chunkSize {
			goodChunks = append(goodChunks, chunk)
			continue
		}

		if len(goodChunks) > 0 {
			mergedText := r.merge(goodChunks, newSep)
			resChunks = append(resChunks, mergedText...)
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
		mergedText := r.merge(goodChunks, newSep)
		resChunks = append(resChunks, mergedText...)
	}

	return resChunks
}

// Split splits text into chunks.
func (r *RecursiveCharSplitter) Split(text string) []string {
	return r.split(text, r.seps)
}
