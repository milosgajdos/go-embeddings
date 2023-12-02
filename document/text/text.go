package text

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

const (
	// DefaultSeparator is default text separator.
	// It's intention is to splitt by paragraphs.
	DefaultSeparator = "\n\n"
)

var (
	// DefaultLenFunc is a default string length function.
	// It counts UTF-8 encoded characters aka runes.
	DefaultLenFunc = utf8.RuneCountInString
	// StringBytesLenFunc counts number of bytes in a string.
	// Faster for some documents, but less accurate for multiling.
	StringBytesLenFunc = func(s string) int { return len(s) }
	//  DefaultSeparators are used in RecursiveSplitter.
	// The splitter recursively keeps splitting document
	// using the separators until done.
	DefaultSeparators = []string{"\n\n", "\n", " ", ""}
)

// LenFunc is used for funcs that calculate string lengths.
type LenFunc func(s string) int

// splitText splits the text on separator into chunks
// optionally keeping the separator in the split chunks.
func splitText(text string, sep string, keepSep bool) []string {
	if sep != "" {
		if keepSep {
			var results []string
			re := regexp.MustCompile("(" + sep + ")")
			splits := re.Split(text, -1)
			for i := 1; i < len(splits); i++ {
				// make sure the separator remains in the result split
				// because Go is silly: https://github.com/golang/go/issues/18868
				results = append(results, sep+splits[i])
			}
			results = append([]string{splits[0]}, results...)
			return results
		}
		re := regexp.MustCompile(sep)
		return re.Split(text, -1)
	}
	// If separator is empty, split the text into individual characters.
	return strings.Split(text, "")
}

// joinChunks joins chunks over sep into a single string
// optionally trimming empty space characters and returns it.
func joinChunks(chunks []string, sep string, trimSpace bool) string {
	text := strings.Join(chunks, sep)
	if trimSpace {
		text = strings.TrimSpace(text)
	}
	return text
}

// boolToInt raturn 1 if b is true
// and false otherwise.
func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
