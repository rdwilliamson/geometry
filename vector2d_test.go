package geometry

import (
	"testing"
)

func TestScalarProjectionOnto(t *testing.T) {
	v1 := Vector2D{2, 3}
	v2 := Vector2D{2, 1}
	want := 7.0 / 5.0
	got := v1.ScalarProjectionOnto(v2)
	if !FuzzyEqual(want, got) {
		t.Error("Vector2D.ScalarProjectionOnto: wanted", want, "got", got)
	}
}

func TestVectorProjectionOnto(t *testing.T) {
	v1 := Vector2D{2, 3}
	v2 := Vector2D{2, 1}
	want := Vector2D{2.8, 1.4}
	got := v1.VectorProjectionOnto(v2)
	if !want.FuzzyEqual(got) {
		t.Error("Vector2D.VectorProjectionOnto: wanted", want, "got", got)
	}
}
