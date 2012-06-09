package geometry

import (
	"testing"
)

func TestFuzzyEqual(t *testing.T) {
	n1 := 1.0
	n2 := n1 + 0.000000000001
	n3 := n1 + 0.0000000000001
	if n1 == n2 {
		t.Fatal("geometry.FuzzyEqual: n1 == n1 + 0.000000000001")
	}
	if n1 == n3 {
		t.Fatal("geometry.FuzzyEqual: n1 == n1 + 0.0000000000001")
	}
	want := false
	got := FuzzyEqual(n1, n2)
	if want != got {
		t.Error("geometry.FuzzyEqual: wanted", want, "got", got)
	}
	want = true
	got = FuzzyEqual(n1, n3)
	if want != got {
		t.Error("geometry.FuzzyEqual: wanted", want, "got", got)
	}
}
