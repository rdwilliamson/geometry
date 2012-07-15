package geometry

import (
	"math"
	"testing"
)

func TestPointDistance3D(t *testing.T) {
	p1 := Point3D{1, 7, -3}
	p2 := Point3D{-5, 4, -2}
	if p1.DistanceTo(p2) != math.Sqrt(46) {
		t.Error("Point3D.DistanceTo")
	}
}

func TestPointDistanceSquared3D(t *testing.T) {
	p1 := Point3D{1, 7, -3}
	p2 := Point3D{-5, 4, -2}
	if p1.SquaredDistanceTo(p2) != 46 {
		t.Error("Point3D.SquaredDistanceTo")
	}
}

func TestPoint3DCore(t *testing.T) {
	p := Point3D{0, 0, 0}
	p = p.Plus(Point3D{1, 2, 3})
	if !p.Equal(Point3D{1, 2, 3}) {
		t.Error("Point3D.Plus")
	}
	p.Add(Point3D{3, 2, 1})
	if !p.Equal(Point3D{4, 4, 4}) {
		t.Error("Point3D.Add")
	}
	p = p.Minus(Point3D{1, 1, 1})
	if !p.Equal(Point3D{3, 3, 3}) {
		t.Error("Point3D.Minus")
	}
	p.Subtract(Point3D{1, 1, 1})
	if !p.Equal(Point3D{2, 2, 2}) {
		t.Error("Point3D.Subtract")
	}
}

func TestPoint3DFuzzyEqual(t *testing.T) {
	p1 := Point3D{1.0, 1.0, 1.0}
	p2 := p1
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
