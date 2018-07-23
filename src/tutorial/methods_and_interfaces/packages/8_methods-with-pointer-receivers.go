package methods_and_interfaces

import (
	"math"
	"fmt"
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

//Here are two reasons to use a pointer receiver.
//The first is so that the method can modify the value that its receiver points to.
//The second is to avoid copying the value on each method call. This can be more efficient
//if the receiver is a large struct, for example.
func MethodPointer8() {
	fmt.Printf("Step 8\n")
	v := &Vertex8{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs8())
	v.Scale8(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs8())
}