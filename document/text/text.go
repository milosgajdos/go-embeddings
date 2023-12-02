package text

import (
	"unicode/utf8"
)

const (
	// DefaultSeparator is default text separator.
	// It's intention is to splitt by paragraphs.
	DefaultSeparator = "\n\n"
	// DefaultChunkSize is default chunk size.
	DefaultChunkSize = 1
	// DefaultChunkOverlap is default chunk overlap.
	DefaultChunkOverlap = 0
)

var (
	// DefaultLenFunc is a default string length function.
	// It counts UTF-8 encoded characters aka runes.
	DefaultLenFunc = utf8.RuneCountInString
	// StringBytesLenFunc counts number of bytes in a string.
	// Faster for some documents, but less accurate for multiling.
	StringBytesLenFunc = func(s string) int { return len(s) }
	// DefaultSeparators are used in RecursiveSplitter.
	// The splitter recursively keeps splitting document
	// using the separators until done.
	DefaultSeparators = []string{"\n\n", "\n", " ", ""}
)

// LenFunc is used for funcs that calculate string lengths.
type LenFunc func(s string) int
