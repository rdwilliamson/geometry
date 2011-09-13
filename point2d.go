package geometry

import (
	"fmt"
)

type Point2D struct {
	X, Y float64
}

func (p Point2D) String() string {
	return fmt.Sprintf("(%g, %g)", p.X, p.Y)
}
