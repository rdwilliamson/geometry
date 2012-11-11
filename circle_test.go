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
	// vertical line resulting in infinite slope
	{Vector2D{0, 1}, Vector2D{math.Sqrt2, math.Sqrt2}, Vector2D{1, 0}, Circle{Vector2D{}, 1}},
}

func testCircleFromThreePoints(d circleFromThreePointsData, t *testing.T) {
	var c Circle
	if c.FromThreePoints(&d.p1, &d.p2, &d.p3); !c.Equal(&d.c) {
		t.Error("Circle.FromThreePoints", d.p1, d.p2, d.p3, d.c, "got", c)
	}
}

func TestCircle2DFromThreePoints(t *testing.T) {
	for _, v := range circleFromThreePointsValues {
		testCircleFromThreePoints(v, t)
	}
}
