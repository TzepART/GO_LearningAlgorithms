package methods

import (
	"math"
)

type Vertex8 struct {
	X, Y float64
}

func (v *Vertex8) Scale8(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex8) Abs8() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}