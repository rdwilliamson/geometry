package geometry

import (
	"math"
	"testing"
)

func TestPoint3DCore(t *testing.T) {
	p := Point3D{0, 0, 0}
	p.Add(&Point3D{3, 2, 1})
	if !p.Equal(&Point3D{3, 2, 1}) {
		t.Error("Point3D.Add")
	}
	p.Subtract(&Point3D{1, 1, 1})
	if !p.Equal(&Point3D{2, 1, 0}) {
		t.Error("Point3D.Subtract")
	}
}

func Benchmark_Point3D_Add(b *testing.B) {
	p1 := &Point3D{0, 1, 2}
	p2 := &Point3D{3, 4, 5}
	for i := 0; i < b.N; i++ {
		p1.Add(p2)
	}
}

func Benchmark_Point3D_Subtract(b *testing.B) {
	p1 := &Point3D{0, 1, 2}
	p2 := &Point3D{3, 4, 5}
	for i := 0; i < b.N; i++ {
		p1.Subtract(p2)
	}
}

func TestPoint3DFuzzyEqual(t *testing.T) {
	p1 := &Point3D{1.0, 1.0, 1.0}
	p2 := &Point3D{1.0, 1.0, 1.0}
	p2.X += 0.0000000000001
	if p1.Equal(p2) {
		t.Error("Point3D.Equal")
	}
	if !p1.FuzzyEqual(p2) {
		t.Error("Point3D.FuzzyEqual")
	}
	p2.Y += 0.000000000001
	if p1.Equal(p2) {
		t.Error("Point3D.Equal")
	}
	if p1.FuzzyEqual(p2) {
		t.Error("Point3D.FuzzyEqual")
	}
}

func Benchmark_Point3D_Equal_1(b *testing.B) {
	p1 := &Point3D{0, 1, 2}
	p2 := &Point3D{0, 1, 2}
	for i := 0; i < b.N; i++ {
		p1.Equal(p2)
	}
}

func Benchmark_Point3D_Equal_2(b *testing.B) {
	p1 := &Point3D{0, 1, 2}
	p2 := &Point3D{0, 1, 3}
	for i := 0; i < b.N; i++ {
		p1.Equal(p2)
	}
}

func Benchmark_Point3D_Equal_3(b *testing.B) {
	p1 := &Point3D{0, 1, 2}
	p2 := &Point3D{0, 3, 3}
	for i := 0; i < b.N; i++ {
		p1.Equal(p2)
	}
}

func Benchmark_Point3D_FuzzyEqual(b *testing.B) {
	p1 := &Point3D{0, 1, 2}
	p2 := &Point3D{0, 1, 2}
	for i := 0; i < b.N; i++ {
		p1.FuzzyEqual(p2)
	}
}

func TestPointDistance3D(t *testing.T) {
	p1 := &Point3D{1, 7, -3}
	p2 := &Point3D{-5, 4, -2}
	if p1.DistanceTo(p2) != math.Sqrt(46) {
		t.Error("Point3D.DistanceTo")
	}
}

func Benchmark_Point3D_Distance(b *testing.B) {
	p1 := &Point3D{0, 1, 2}
	p2 := &Point3D{1, 2, 3}
	for i := 0; i < b.N; i++ {
		p1.DistanceTo(p2)
	}
}

func TestPointDistanceSquared3D(t *testing.T) {
	p1 := &Point3D{1, 7, -3}
	p2 := &Point3D{-5, 4, -2}
	if p1.SquaredDistanceTo(p2) != 46 {
		t.Error("Point3D.SquaredDistanceTo")
	}
}

func Benchmark_Point3D_SquaredDistanceTo(b *testing.B) {
	p1 := &Point3D{0, 1, 2}
	p2 := &Point3D{1, 2, 3}
	for i := 0; i < b.N; i++ {
		p1.SquaredDistanceTo(p2)
	}
}
