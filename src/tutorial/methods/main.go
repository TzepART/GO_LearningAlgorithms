package main

import (
	"tutorial/methods/packages"
	"fmt"
	"math"
)

// https://tour.golang.org/methods/1
func main() {
	//1. Use method is a function with a special receiver argument.
	v := packages.Vertex{3, 4}
	fmt.Println(v.Abs())

	//2. Use method is a function without a special receiver argument.
	v2 := packages.Vertex2{4, 5}
	fmt.Println(packages.Abs2(v2))

	//3. Declare a method on non-struct types
	f := packages.MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs3())

	//4. Declare methods with pointer receivers
	v4 := packages.Vertex4{3, 4}
	v4.Scale(10)
	fmt.Println(v4.Abs4())
}
