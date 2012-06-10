package geometry

import (
	"testing"
)

func TestConversions(t *testing.T) {
	_ = Vector2D(Point2D{})
	_ = Point2D(Vector2D{})
	_ = Vector3D(Point3D{})
	_ = Point3D(Vector3D{})
}

func TestFuzzyEqual(t *testing.T) {
	n1 := 1.0
	n2 := n1 + 0.000000000001
	n3 := n1 + 0.0000000000001
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
