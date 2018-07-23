package methods_and_interfaces

import (
	"math"
	"fmt"
)

type Vertex4 struct {
	X, Y float64
}

func (v Vertex4) Abs4() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex4) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func MethodPointer4() {
	v4 := Vertex4{3, 4}
	v4.Scale(10)
	fmt.Println("Step 4 -", v4.Abs4())
}
