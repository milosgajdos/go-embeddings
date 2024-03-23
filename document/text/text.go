package text

import (
	"unicode/utf8"
)

const (
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
	// DefaultSeparator is default text separator.
	// Its intention is to splitt by paragraphs.
	DefaultSeparator = Sep{Value: "\n\n"}
	// DefaultSeparators are used in RecursiveCharSplitter.
	// RecursiveCharSplitter keeps splitting document
	// recursively using the separators until done.
	DefaultSeparators = []Sep{
		{Value: "\n\n"},
		{Value: "\n"},
		{Value: " "},
		{Value: ""},
	}
)

// Sep is a text separator.
type Sep struct {
	Value    string
	IsRegexp bool
}

// LenFunc is used for funcs that calculate string lengths.
type LenFunc func(s string) int
