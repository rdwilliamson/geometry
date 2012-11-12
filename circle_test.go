package geometry

import (
	"math"
	"testing"
)

type circleFromThreePointsData struct {
	p1, p2, p3 Vector2D
	c          Circle
}

var circleFromThreePointsValues = []circleFromThreePointsData{
	// two points form a vertical line (infinite slope)
	// {Vector2D{0, 1}, Vector2D{1, 1}, Vector2D{1, 0}, Circle{Vector2D{}, 1}},
	// colinear
	{Vector2D{0, 1}, Vector2D{0, 2}, Vector2D{0, 3}, Circle{Vector2D{math.NaN(), math.NaN()}, math.NaN()}},
	// regular
	{Vector2D{0, 1}, Vector2D{math.Sqrt2 / 2, math.Sqrt2 / 2}, Vector2D{1, 0}, Circle{Vector2D{}, 1}},
}

func testCircleFromThreePoints(d circleFromThreePointsData, t *testing.T) {
	var c Circle
	c.FromThreePoints(&d.p1, &d.p2, &d.p3)
	if !c.C.nanFuzzyEqual(&d.c.C) || !nanFuzzyEqual(c.R, d.c.R) {
		t.Error("Circle.FromThreePoints", d.p1, d.p2, d.p3, d.c, "got", c)
	}
}

func TestCircleFromThreePoints(t *testing.T) {
	for _, v := range circleFromThreePointsValues {
		testCircleFromThreePoints(v, t)
	}
}

func Benchmark_Circle_FromThreePoints(b *testing.B) {
	var c Circle
	p1, p2, p3 := Vector2D{1, 2}, Vector2D{3, 4}, Vector2D{5, 6}
	for i := 0; i < b.N; i++ {
		c.FromThreePoints(&p1, &p2, &p3)
	}
}
