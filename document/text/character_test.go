package text

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCharSplitter(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		size    int
		overlap int
		trim    bool
		keepSep bool
		sep     Sep
		input   string
		exp     []string
	}{
		{
			size:    7,
			overlap: 3,
			sep:     Sep{Value: " "},
			input:   "foo bar baz 123",
			exp:     []string{"foo bar", "bar baz", "baz 123"},
		},
		{
			size:    2,
			overlap: 0,
			sep:     Sep{Value: " "},
			input:   "foo  bar",
			exp:     []string{"foo", "bar"},
		},
		{
			size:    3,
			overlap: 1,
			sep:     Sep{Value: " "},
			input:   "foo bar baz a a",
			exp:     []string{"foo", "bar", "baz", "a a"},
		},
		{
			size:    3,
			overlap: 1,
			sep:     Sep{Value: " "},
			input:   "a a foo bar baz",
			exp:     []string{"a a", "foo", "bar", "baz"},
		},
		{
			size:    1,
			overlap: 1,
			sep:     Sep{Value: " "},
			input:   "foo bar baz 123",
			exp:     []string{"foo", "bar", "baz", "123"},
		},
		{
			size:    1,
			overlap: 0,
			keepSep: true,
			sep:     Sep{Value: ".", IsRegexp: false},
			input:   "foo.bar.baz.123",
			exp:     []string{"foo", ".bar", ".baz", ".123"},
		},
		{
			size:    1,
			overlap: 0,
			sep:     Sep{Value: ".", IsRegexp: false},
			input:   "foo.bar.baz.123",
			exp:     []string{"foo", "bar", "baz", "123"},
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
		cs := NewCharSplitter().
			WithSplitter(s).
			WithSep(tc.sep)

		t.Run(fmt.Sprintf("sep=%#v,size=%d,overlap=%d,trim=%v,keepSep=%v",
			tc.sep, tc.size, tc.overlap, tc.trim, tc.keepSep),
			func(t *testing.T) {
				t.Parallel()
				splits := cs.Split(tc.input)
				if !reflect.DeepEqual(splits, tc.exp) {
					t.Errorf("expected: %#v, got: %#v", tc.exp, splits)
				}
			})
	}
}
