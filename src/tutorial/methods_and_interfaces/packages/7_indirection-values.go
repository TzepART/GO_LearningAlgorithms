package methods_and_interfaces

import (
	"math"
	"fmt"
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

func MethodInderection7() {
	v7 := Vertex7{3, 4}
	fmt.Println("Step 7 -", v7.Abs7())
	fmt.Println("Step 7 -", AbsFunc7(v7))
	//by pointer result similar
	p7 := &Vertex7{4, 3}
	fmt.Println("Step 7 -", p7.Abs7())
	fmt.Println("Step 7 -", AbsFunc7(*p7))
}