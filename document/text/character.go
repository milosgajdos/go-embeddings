package text

import (
	"regexp"
)

// CharSplitter is a character text splitter.
// It splits texts into chunks over
// a separator which is either a string
// or a regular expression.
type CharSplitter struct {
	*Splitter
	sep        string
	isSepRegex bool
}

// NewSplitter creates a new splitter
// with default options and returns it.
func NewCharSplitter() *CharSplitter {
	return &CharSplitter{
		Splitter: NewSplitter(),
		sep:      DefaultSeparator,
	}
}

// WithSplitter sets the splitter
func (s *CharSplitter) WithSplitter(splitter *Splitter) *CharSplitter {
	s.Splitter = splitter
	return s
}

// WithSep sets the separator.
func (s *CharSplitter) WithSep(sep string, isSepRegex bool) *CharSplitter {
	s.sep = sep
	s.isSepRegex = isSepRegex
	return nil
}

// Split splits text into chunks.
func (s *CharSplitter) Split(text string) []string {
	sep := s.sep
	if !s.isSepRegex {
		sep = regexp.QuoteMeta(s.sep)
	}
	chunks := s.splitText(text, sep)
	sep = ""
	if !s.keepSep {
		sep = s.sep
	}
	return s.merge(chunks, sep)
}
