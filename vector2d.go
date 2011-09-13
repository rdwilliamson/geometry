package geometry

import (
	"fmt"
	"math"
)

type Vector2D struct {
	X, Y float64
}

func (v *Vector2D) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vector2D) String() string {
	return fmt.Sprintf("(%g, %g)", v.X, v.Y)
}

func DotProduct2D(v1, v2 *Vector2D) float64 {
	return v1.X*v2.X + v1.Y*v2.Y
}
