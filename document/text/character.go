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
	sep Sep
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
func (s *CharSplitter) WithSep(sep Sep) *CharSplitter {
	s.sep = sep
	return s
}

// Split splits text into chunks.
func (s *CharSplitter) Split(text string) []string {
	sep := Sep{Value: s.sep.Value, IsRegexp: s.sep.IsRegexp}
	if !sep.IsRegexp {
		sep.Value = regexp.QuoteMeta(sep.Value)
	}
	splits := s.splitText(text, sep)
	if s.keepSep {
		sep.Value = ""
	}
	return s.merge(splits, sep)
}
