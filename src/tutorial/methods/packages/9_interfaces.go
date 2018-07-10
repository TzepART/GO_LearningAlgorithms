package methods

import (
	"math"
)

type Abser interface {
	Abs9() float64
}

type MyFloat9 float64

func (f MyFloat9) Abs9() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex9 struct {
	X, Y float64
}

func (v *Vertex9) Abs9() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}