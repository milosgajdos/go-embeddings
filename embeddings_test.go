package embeddings

import "testing"

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
