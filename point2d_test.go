package geometry

import (
	"testing"
)

func TestPoint2DCore(t *testing.T) {
	p := &Point2D{0, 0}
	p.Add(&Point2D{2, 1})
	if !p.Equal(&Point2D{2, 1}) {
		t.Error("Point2D.Add")
	}
	p.Subtract(&Point2D{1, 1})
	if !p.Equal(&Point2D{1, 0}) {
		t.Error("Point2D.Subtract")
	}
}

func Benchmark_Point2D_Add(b *testing.B) {
	p1, p2 := &Point2D{1, 1}, &Point2D{2, 1}
	for i := 0; i < b.N; i++ {
		p1.Add(p2)
	}
}

func Benchmark_Point2D_Subtract(b *testing.B) {
	p1, p2 := &Point2D{1, 1}, &Point2D{2, 1}
	for i := 0; i < b.N; i++ {
		p1.Subtract(p2)
	}
}

func TestPointDistance2D(t *testing.T) {
	p1 := &Point2D{-2, 1}
	p2 := &Point2D{1, 5}
	if p1.DistanceTo(p2) != 5 {
		t.Error("Point2D.DistanceTo")
	}
}

func Benchmark_Point2D_Distance(b *testing.B) {
	p1, p2 := &Point2D{1, 1}, &Point2D{2, 1}
	for i := 0; i < b.N; i++ {
		p1.DistanceTo(p2)
	}
}

func TestPointDistanceSquared2D(t *testing.T) {
	p1 := &Point2D{-2, 1}
	p2 := &Point2D{1, 5}
	if p1.SquaredDistanceTo(p2) != 25 {
		t.Error("Point2D.SquaredDistanceTo")
	}
}

func Benchmark_Point2D_SquaredDistance(b *testing.B) {
	p1, p2 := &Point2D{1, 1}, &Point2D{2, 1}
	for i := 0; i < b.N; i++ {
		p1.SquaredDistanceTo(p2)
	}
}

func TestPoint2DFuzzyEqual(t *testing.T) {
	p1 := &Point2D{1.0, 1.0}
	p2 := &Point2D{1.0, 1.0}
	if !p1.Equal(p2) {
		t.Error("Point2D.Equal")
	}
	p2.X += 0.0000000000001
	if p1.Equal(p2) {
		t.Error("Point2D.Equal")
	}
	if !p1.FuzzyEqual(p2) {
		t.Error("Point2D.FuzzyEqual")
	}
	p2.Y += 0.000000000001
	if p1.Equal(p2) {
		t.Error("Point2D.Equal")
	}
	if p1.FuzzyEqual(p2) {
		t.Error("Point2D.FuzzyEqual")
	}
}

func Benchmark_Point2D_Equal_1(b *testing.B) {
	// x values are not equal so no need to check y
	p1, p2 := &Point2D{1, 1}, &Point2D{2, 1}
	for i := 0; i < b.N; i++ {
		p1.Equal(p2)
	}
}

func Benchmark_Point2D_Equal_2(b *testing.B) {
	p1, p2 := &Point2D{1, 1}, &Point2D{1, 1}
	for i := 0; i < b.N; i++ {
		p1.Equal(p2)
	}
}

func Benchmark_Point2D_FuzzyEqual(b *testing.B) {
	p1, p2 := &Point2D{1, 1}, &Point2D{2, 3}
	for i := 0; i < b.N; i++ {
		p1.FuzzyEqual(p2)
	}
}
