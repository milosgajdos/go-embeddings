package text

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRecursiveCharSplitter(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		size    int
		overlap int
		trim    bool
		keepSep bool
		seps    []Sep
		input   string
		exp     []string
	}{
		{
			size:    10,
			overlap: 1,
			trim:    true,
			keepSep: true,
			seps:    DefaultSeparators,
			input: `Hi.` + "\n\n" + `I'm Harrison.` + "\n\n" + `How? Are? You?` + "\n" + `Okay then f f f f.
This is a weird text to write, but gotta test the splittingggg some how.

Bye!` + "\n\n" + `-H.`,
			exp: []string{
				"Hi.",
				"I'm",
				"Harrison.",
				"How? Are?",
				"You?",
				"Okay then",
				"f f f f.",
				"This is a",
				"weird",
				"text to",
				"write,",
				"but gotta",
				"test the",
				"splitting",
				"gggg",
				"some how.",
				"Bye!",
				"-H.",
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		s := NewSplitterWithConfig(Config{
			ChunkSize:    tc.size,
			ChunkOverlap: tc.overlap,
			TrimSpace:    tc.trim,
			KeepSep:      tc.keepSep,
			LenFunc:      DefaultLenFunc,
		})
		cs := NewRecursiveCharSplitter().
			WithSplitter(s).
			WithSeps(tc.seps)

		t.Run(fmt.Sprintf("sep=%#v,size=%d,overlap=%d,trim=%v,keepSep=%v",
			tc.seps, tc.size, tc.overlap, tc.trim, tc.keepSep),
			func(t *testing.T) {
				t.Parallel()
				splits := cs.Split(tc.input)
				if !reflect.DeepEqual(splits, tc.exp) {
					t.Errorf("expected: %#v, got: %#v", tc.exp, splits)
				}
			})
	}
}
