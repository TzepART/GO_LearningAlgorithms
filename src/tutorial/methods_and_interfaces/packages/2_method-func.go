package methods_and_interfaces

import (
	"math"
	"fmt"
)

type Vertex2 struct {
	X, Y float64
}

func Abs2(v Vertex2) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func MethodExample2() {
	v2 := Vertex2{4, 5}
	fmt.Println("Step 2 -", Abs2(v2))
}
