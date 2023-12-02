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
	seps       []string
	isSepRegex bool
}

// NewSplitter creates a new splitter and returns it.
func NewRecursiveCharSplitter() *RecursiveCharSplitter {
	return &RecursiveCharSplitter{
		Splitter: NewSplitter(),
		seps:     DefaultSeparators,
	}
}

// WithSplitter sets the splitter
func (r *RecursiveCharSplitter) WithSplitter(splitter *Splitter) *RecursiveCharSplitter {
	r.Splitter = splitter
	return r
}

// WithSeps sets separators
func (r *RecursiveCharSplitter) WithSeps(seps []string, isSepRegex bool) *RecursiveCharSplitter {
	r.seps = seps
	r.isSepRegex = isSepRegex
	return nil
}

func (r *RecursiveCharSplitter) split(text string, seps []string) []string {
	var (
		resChunks []string
		newSeps   []string
	)

	sep := seps[len(seps)-1]

	for i, s := range seps {
		if !r.isSepRegex {
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
	if !r.isSepRegex {
		newSep = regexp.QuoteMeta(sep)
	}
	chunks := r.splitText(text, newSep)

	var goodChunks []string

	if r.keepSep {
		newSep = ""
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
