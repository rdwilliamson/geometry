package geometry

import (
	"testing"
	"math"
)

func TestLine3DToVector(t *testing.T) {
	l := Line3D{Point3D{1, 1, 1}, Point3D{2, 3, 4}}
	if v := l.ToVector(); !v.Equal(&Vector3D{1, 2, 3}) {
		t.Error("Line3D.ToVector")
	}
}

func Benchmark_Line3D_ToVector(b *testing.B) {
	l := Line3D{Point3D{1, 1, 1}, Point3D{2, 3, 4}}
	for i := 0; i < b.N; i++ {
		l.ToVector()
	}
}

func TestLine3DLength(t *testing.T) {
	l := Line3D{Point3D{-5, 4, 0}, Point3D{1, 1, 1}}
	if l.Length() != math.Sqrt(46) {
		t.Error("Line3D.Length")
	}
}

func Benchmark_Line3D_Length(b *testing.B) {
	l := Line3D{Point3D{1, 1, 1}, Point3D{2, 3, 4}}
	for i := 0; i < b.N; i++ {
		l.Length()
	}
}

func TestLine3DLengthSquared(t *testing.T) {
	l := Line3D{Point3D{-5, 4, 0}, Point3D{1, 1, 1}}
	if l.LengthSquared() != 46 {
		t.Error("Line3D.LengthSquared")
	}
}

func Benchmark_Line3D_LengthSquared(b *testing.B) {
	l := Line3D{Point3D{1, 1, 1}, Point3D{2, 3, 4}}
	for i := 0; i < b.N; i++ {
		l.LengthSquared()
	}
}

func TestLine3DMidpoint(t *testing.T) {
	l := Line3D{Point3D{-5, 4, 0}, Point3D{1, 1, 1}}
	if p := l.Midpoint(); !p.Equal(&Point3D{3, -1.5, 0.5}) {
		t.Error("Line3D.Midpoint")
	}
}

func Benchmark_Line3D_Midpoint(b *testing.B) {
	l := Line3D{Point3D{1, 1, 1}, Point3D{2, 3, 4}}
	for i := 0; i < b.N; i++ {
		l.Midpoint()
	}
}

func TestLine3DEqual(t *testing.T) {
	l1 := &Line3D{Point3D{1, 2, 3}, Point3D{1, 1, 1}}
	l2 := &Line3D{Point3D{1, 2, 3}, Point3D{1, 1, 1}}
	if !l1.Equal(l2) {
		t.Error("Line3D.Equal")
	}
}

func Benchmark_Line3D_Equal(b *testing.B) {
	l1 := &Line3D{Point3D{1, 2, 3}, Point3D{1, 1, 1}}
	l2 := &Line3D{Point3D{1, 2, 3}, Point3D{1, 1, 3}}
	for i := 0; i < b.N; i++ {
		l1.Equal(l2)
	}
}

func TestLine3DFuzzyFuzzyEqual(t *testing.T) {
	l1 := &Line3D{Point3D{1, 2, 3}, Point3D{1, 1, 1}}
	l2 := &Line3D{Point3D{1, 1, 1}, Point3D{1, 2, 3}}
	l1.P1.X += 0.0000000000001
	if !l1.FuzzyEqual(l2) || !l1.FuzzyEqual(l1) {
		t.Error("Line3D.FuzzyEqual")
	}
	l2.P1.X += 0.000000000001
	if l1.FuzzyEqual(l2) || !l1.FuzzyEqual(l1) {
		t.Error("Line3D.FuzzyEqual")
	}
}

func Benchmark_Line3D_FuzzyEqual(b *testing.B) {
	l1 := &Line3D{Point3D{1, 2, 3}, Point3D{1, 1, 1}}
	l2 := &Line3D{Point3D{1, 1, 1}, Point3D{1, 2, 3}}
	for i := 0; i < b.N; i++ {
		l1.FuzzyEqual(l2)
	}
}
