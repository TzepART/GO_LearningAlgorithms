package main

import (
	"tutorial/methods/packages"
	"fmt"
	"math"
)

func main() {
	//Use method is a function with a special receiver argument.
	v := packages.Vertex{3, 4}
	fmt.Println(v.Abs())

	//Use method is a function without a special receiver argument.
	v2 := packages.Vertex2{4, 5}
	fmt.Println(packages.Abs2(v2))

	//declare a method on non-struct types
	f := packages.MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs3())
}
