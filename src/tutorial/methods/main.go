package main

import (
	"tutorial/methods/packages"
	"fmt"
	"math"
)

// https://tour.golang.org/methods/1
func main() {
	//1. Use method is a function with a special receiver argument.
	v := methods.Vertex{3, 4}
	fmt.Println("Step 1 -",v.Abs())

	//2. Use method is a function without a special receiver argument.
	v2 := methods.Vertex2{4, 5}
	fmt.Println("Step 2 -",methods.Abs2(v2))

	//3. Declare a method on non-struct types
	f := methods.MyFloat(-math.Sqrt2)
	fmt.Println("Step 3 -",f.Abs3())

	//4. Declare methods with pointer receivers
	v4 := methods.Vertex4{3, 4}
	v4.Scale(10)
	fmt.Println("Step 4 -",v4.Abs4())

	//5. Here we see the Abs and Scale methods rewritten as functions
	v5 := methods.Vertex5{3, 4}
	methods.Scale5(&v5, 5)
	fmt.Println("Step 5 -",methods.Abs5(v5))

	//6. Methods and pointer indirection
	v6 := methods.Vertex6{3, 4}
	v6.Scale6(2)
	methods.ScaleFunc6(&v6, 10)

	p := &methods.Vertex6{4, 3}
	p.Scale6(3)
	methods.ScaleFunc6(p, 8)

	fmt.Println("Step 6 -",v6, p)
}
