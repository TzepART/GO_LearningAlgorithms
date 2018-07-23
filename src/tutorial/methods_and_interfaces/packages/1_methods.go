package methods_and_interfaces

import (
	"math"
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func MethodExample1()  {
	v := Vertex{3, 4}
	fmt.Println("Step 1 -",v.Abs())
}

