package methods_and_interfaces

import (
	"math"
)

type Vertex5 struct {
	X, Y float64
}

func Abs5(v Vertex5) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale5(v *Vertex5, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
