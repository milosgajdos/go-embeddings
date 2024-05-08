package embeddings

import (
	"reflect"
	"testing"
)

func TestToFloat32(t *testing.T) {
	e := Embedding{
		Vector: []float64{1.0, 2.0, 3.0},
	}

	exp := []float32{1.0, 2.0, 3.0}
	got := e.ToFloat32()

	if len(got) != len(exp) {
		t.Fatalf("expected %d vals, got %v", len(exp), len(got))
	}

	for i, f := range got {
		if exp[i] != f {
			t.Fatalf("expected %v, got %v", exp, got)
		}
	}
}

func TestBase64Decode(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name    string
		given   string
		exp     []float64
		wantErr bool
	}{
		{
			name:    "valid",
			given:   "H4XrUbgeCUAAAAAAAABFwESLbOf7QI9A",
			exp:     []float64{3.14, -42.0, 1000.123},
			wantErr: false,
		},
		{
			name:    "invalid",
			given:   "garbage",
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			embBase64 := Base64(tc.given)
			got, err := embBase64.Decode()
			if err != nil {
				if !tc.wantErr {
					t.Fatalf("unexpected error: %v", err)
				}
				return
			}
			// no error, but we expect one
			if tc.wantErr {
				t.Fatal("expected error")
			}
			if !reflect.DeepEqual(got.Vector, tc.exp) {
				t.Fatalf("expected: %v, got: %v", tc.exp, got.Vector)
			}
		})
	}
}
