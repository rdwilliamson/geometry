package geometry

import (
	"math"
	"testing"
)

func TestGeneral(t *testing.T) {
	nan32 := float32(math.NaN())
	if !math.IsNaN(float64(nan32)) {
		t.Log("math.IsNaN(float32(float64(math.NaN()))) is false!")
	}
}

func TestAbs32(t *testing.T) {
	n1 := float32(-1.0)
	n2 := float32(-0.0)
	n3 := float32(1.0)
	n4 := float32(math.Inf(1))
	n5 := float32(math.Inf(-1))
	n6 := float32(math.NaN())
	want := float32(math.Abs(float64(n1)))
	got := Abs32(n1)
	if got != want {
		t.Error("geometry.Abs32: wanted", want, "got", got)
	}
	want = float32(math.Abs(float64(n2)))
	got = Abs32(n2)
	if got != want {
		t.Error("geometry.Abs32: wanted", want, "got", got)
	}
	want = float32(math.Abs(float64(n3)))
	got = Abs32(n3)
	if got != want {
		t.Error("geometry.Abs32: wanted", want, "got", got)
	}
	want = float32(math.Abs(float64(n4)))
	got = Abs32(n4)
	if got != want {
		t.Error("geometry.Abs32: wanted", want, "got", got)
	}
	want = float32(math.Abs(float64(n5)))
	got = Abs32(n5)
	if got != want {
		t.Error("geometry.Abs32: wanted", want, "got", got)
	}
	want = float32(math.Abs(float64(n6)))
	got = Abs32(n6)
	if !math.IsNaN(float64(got)) {
		t.Error("geometry.Abs32: wanted", want, "got", got)
	}
}

func TestMin32(t *testing.T) {
	n1 := float32(0.0)
	n2 := float32(1.0)
	n3 := float32(math.Inf(-1))
	n4 := float32(math.NaN())
	n5 := float32(-0.0)
	want := float32(math.Min(float64(n1), float64(n2)))
	got := Min32(n1, n2)
	if got != want {
		t.Error("geometry.Min32: wanted", want, "got", got)
	}
	want = float32(math.Min(float64(n2), float64(n3)))
	got = Min32(n2, n3)
	if got != want {
		t.Error("geometry.Min32: wanted", want, "got", got)
	}
	want = float32(math.Min(float64(n3), float64(n4)))
	got = Min32(n3, n4)
	if got != want {
		t.Error("geometry.Min32: wanted", want, "got", got)
	}
	want = float32(math.Min(float64(n1), float64(n5)))
	got = Min32(n1, n5)
	if got != want {
		t.Error("geometry.Min32: wanted", want, "got", got)
	}
}
