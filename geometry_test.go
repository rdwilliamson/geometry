package geometry

import "testing"

func TestFuzzyEqual(t *testing.T) {
	n1 := 1.0
	n2 := n1 + 1e-12
	n3 := n1 + 1e-13
	if n1 == n2 {
		t.Fatal("geometry.FuzzyEqual")
	}
	if n1 == n3 {
		t.Fatal("geometry.FuzzyEqual")
	}
	if FuzzyEqual(n1, n2) {
		t.Error("geometry.FuzzyEqual")
	}
	if !FuzzyEqual(n1, n3) {
		t.Error("geometry.FuzzyEqual")
	}
}

func Benchmark_FuzzyEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FuzzyEqual(1.0, 1+1e-13)
	}
}
