package main

import (
	"tutorial/methods/packages"
	"fmt"
	"math"
)

// https://tour.golang.org/methods/1
func main() {
	step8()
}

//1. Use method is a function with a special receiver argument.
func step1()  {
	v := methods.Vertex{3, 4}
	fmt.Println("Step 1 -",v.Abs())
}

//2. Use method is a function without a special receiver argument.
func step2() {
	v2 := methods.Vertex2{4, 5}
	fmt.Println("Step 2 -", methods.Abs2(v2))
}

//3. Declare a method on non-struct types
func step3() {
	f := methods.MyFloat(-math.Sqrt2)
	fmt.Println("Step 3 -", f.Abs3())
}

//4. Declare methods with pointer receivers
func step4() {
	v4 := methods.Vertex4{3, 4}
	v4.Scale(10)
	fmt.Println("Step 4 -", v4.Abs4())
}

//5. Here we see the Abs and Scale methods rewritten as functions
func step5() {
	v5 := methods.Vertex5{3, 4}
	methods.Scale5(&v5, 5)
	fmt.Println("Step 5 -", methods.Abs5(v5))
}

//6. Methods and pointer indirection
func step6() {
	v6 := methods.Vertex6{3, 4}
	v6.Scale6(2)
	methods.ScaleFunc6(&v6, 10)
	p := &methods.Vertex6{4, 3}
	p.Scale6(3)
	methods.ScaleFunc6(p, 8)
	fmt.Println("Step 6 -", v6, p)
}


//7. The equivalent thing happens in the reverse direction.
func step7() {
	v7 := methods.Vertex7{3, 4}
	fmt.Println("Step 7 -", v7.Abs7())
	fmt.Println("Step 7 -", methods.AbsFunc7(v7))
	//by pointer result similar
	p7 := &methods.Vertex7{4, 3}
	fmt.Println("Step 7 -", p7.Abs7())
	fmt.Println("Step 7 -", methods.AbsFunc7(*p7))
}

//8. Here are two reasons to use a pointer receiver.
//The first is so that the method can modify the value that its receiver points to.
//The second is to avoid copying the value on each method call. This can be more efficient
//if the receiver is a large struct, for example.
func step8() {
	fmt.Printf("Step 8\n")
	v := &methods.Vertex8{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs8())
	v.Scale8(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs8())
}





