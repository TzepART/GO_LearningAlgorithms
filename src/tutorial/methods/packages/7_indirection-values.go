package methods

import (
	"math"
)

type Vertex7 struct {
	X, Y float64
}

func (v Vertex7) Abs7() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc7(v Vertex7) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}