package geometry

import (
	"math"
	"testing"
)

type circleAreaData struct {
	c    Circle
	area float64
}

var circleAreaValues = []circleAreaData{
	{Circle{Vector2D{}, 1}, math.Pi},
}

func testCircleArea(d circleAreaData, t *testing.T) {
	if got := d.c.Area(); got != math.Pi {
		t.Error("Circle.Area", d.c, "got", got)
	}
}

func TestCircleArea(t *testing.T) {
	for _, v := range circleAreaValues {
		testCircleArea(v, t)
	}
}

type circleCopyData struct {
	c Circle
}

var circleCopyValues = []circleCopyData{
	{Circle{Vector2D{1, 2}, 3}},
}

func testCircleCopy(d circleCopyData, t *testing.T) {
	var c Circle
	if c.Equal(&d.c) || !c.Copy(&d.c).Equal(&d.c) {
		t.Error("Circle.Copy")
	}
}

func TestCircleCopy(t *testing.T) {
	for _, v := range circleCopyValues {
		testCircleCopy(v, t)
	}
}

type circleEqualData struct {
	c1, c2 Circle
	equal  bool
}

var circleEqualValues = []circleEqualData{
	{Circle{Vector2D{1, 2}, 3}, Circle{Vector2D{1, 2}, 3}, true},
	{Circle{Vector2D{1, 2}, 3}, Circle{Vector2D{1, 2}, 4}, false},
}

func testCircleEqual(d circleEqualData, t *testing.T) {
	if d.c1.Equal(&d.c2) != d.equal {
		t.Error("Circle.Equal", d.c1, d.c2, d.equal)
	}
}

func TestCircleEqual(t *testing.T) {
	for _, v := range circleEqualValues {
		testCircleEqual(v, t)
	}
}

type circleFromThreePointsData struct {
	p1, p2, p3 Vector2D
	c          Circle
}

var circleFromThreePointsValues = []circleFromThreePointsData{
	// two points form a vertical line
	{Vector2D{0, 1}, Vector2D{0, 0}, Vector2D{1, 0}, Circle{Vector2D{0.5, 0.5}, math.Sqrt2 / 2}},
	{Vector2D{0, 0}, Vector2D{0, 1}, Vector2D{1, 0}, Circle{Vector2D{0.5, 0.5}, math.Sqrt2 / 2}},
	{Vector2D{1, 0}, Vector2D{0, 1}, Vector2D{0, 0}, Circle{Vector2D{0.5, 0.5}, math.Sqrt2 / 2}},
	{Vector2D{1, 0}, Vector2D{0, 0}, Vector2D{0, 1}, Circle{Vector2D{0.5, 0.5}, math.Sqrt2 / 2}},
	// collinear
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

type circleFuzzyEqualData struct {
	c1, c2 Circle
	equal  bool
}

var circleFuzzyEqualValues = []circleFuzzyEqualData{
	{Circle{Vector2D{}, 1}, Circle{Vector2D{}, 1}, true},
	{Circle{Vector2D{1, 1}, 1}, Circle{Vector2D{1 + 1e-13, 1}, 1}, true},
	{Circle{Vector2D{1, 1}, 1}, Circle{Vector2D{1 + 1e-12, 1}, 1}, false},
	{Circle{Vector2D{1, 1}, 1}, Circle{Vector2D{1, 1 + 1e-13}, 1}, true},
	{Circle{Vector2D{1, 1}, 1}, Circle{Vector2D{1, 1 + 1e-12}, 1}, false},
	{Circle{Vector2D{1, 1}, 1}, Circle{Vector2D{1, 1}, 1 + 1e-13}, true},
	{Circle{Vector2D{1, 1}, 1}, Circle{Vector2D{1, 1}, 1 + 1e-12}, false},
}

func testCircleFuzzyEqual(d circleFuzzyEqualData, t *testing.T) {
	if d.c1.FuzzyEqual(&d.c2) != d.equal {
		t.Error("Circle.FuzzyEqual", d.c1, d.c2, d.equal)
	}
}

func TestCircleFuzzyEqual(t *testing.T) {
	for _, v := range circleFuzzyEqualValues {
		testCircleFuzzyEqual(v, t)
	}
}

type circlePerimeterData struct {
	c          Circle
	permimeter float64
}

var circlePermimeterValues = []circlePerimeterData{
	{Circle{Vector2D{}, 1}, 2 * math.Pi},
}

func testCirclePerimeter(d circlePerimeterData, t *testing.T) {
	if d.c.Perimeter() != d.permimeter {
		t.Error("Circle.Perimeter", d.c, d.permimeter)
	}
}

func TestCirclePreimeter(t *testing.T) {
	for _, v := range circlePermimeterValues {
		testCirclePerimeter(v, t)
	}
}
