package methods_and_interfaces

import (
	"math"
	"fmt"
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

func MethodPointer5() {
	v5 := Vertex5{3, 4}
	Scale5(&v5, 5)
	fmt.Println("Step 5 -", Abs5(v5))
}