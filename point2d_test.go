package geometry

import (
	"testing"
)

func TestPointDistance2D(t *testing.T) {
	p1 := Point2D{-2, 1}
	p2 := Point2D{1, 5}
	if p1.DistanceTo(p2) != 5 {
		t.Error("Point2D.DistanceTo")
	}
}

func TestPointDistanceSquared2D(t *testing.T) {
	p1 := Point2D{-2, 1}
	p2 := Point2D{1, 5}
	if p1.SquaredDistanceTo(p2) != 25 {
		t.Error("Point2D.SquaredDistanceTo")
	}
}

func TestPoint2DCore(t *testing.T) {
	p := Point2D{0, 0}
	p = p.Plus(Point2D{1, 2})
	if !p.Equal(Point2D{1, 2}) {
		t.Error("Point2D.Plus")
	}
	p.Add(Point2D{2, 1})
	if !p.Equal(Point2D{3, 3}) {
		t.Error("Point2D.Add")
	}
	p = p.Minus(Point2D{1, 1})
	if !p.Equal(Point2D{2, 2}) {
		t.Error("Point2D.Minus")
	}
	p.Subtract(Point2D{1, 1})
	if !p.Equal(Point2D{1, 1}) {
		t.Error("Point2D.Subtract")
	}
}

func TestPoint2DFuzzyEqual(t *testing.T) {
	p1 := Point2D{1.0, 1.0}
	p2 := p1
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
