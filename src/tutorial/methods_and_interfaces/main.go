package main

import (
	"fmt"
	"math"
	"tutorial/methods_and_interfaces/packages"
)

// https://tour.golang.org/methods_and_interfaces/1
func main() {
	step11()
}

//1. Use method is a function with a special receiver argument.
func step1()  {
	v := methods_and_interfaces.Vertex{3, 4}
	fmt.Println("Step 1 -",v.Abs())
}

//2. Use method is a function without a special receiver argument.
func step2() {
	v2 := methods_and_interfaces.Vertex2{4, 5}
	fmt.Println("Step 2 -", methods_and_interfaces.Abs2(v2))
}

//3. Declare a method on non-struct types
func step3() {
	f := methods_and_interfaces.MyFloat(-math.Sqrt2)
	fmt.Println("Step 3 -", f.Abs3())
}

//4. Declare methods_and_interfaces with pointer receivers
func step4() {
	v4 := methods_and_interfaces.Vertex4{3, 4}
	v4.Scale(10)
	fmt.Println("Step 4 -", v4.Abs4())
}

//5. Here we see the Abs and Scale methods_and_interfaces rewritten as functions
func step5() {
	v5 := methods_and_interfaces.Vertex5{3, 4}
	methods_and_interfaces.Scale5(&v5, 5)
	fmt.Println("Step 5 -", methods_and_interfaces.Abs5(v5))
}

//6. methods_and_interfaces and pointer indirection
func step6() {
	v6 := methods_and_interfaces.Vertex6{3, 4}
	v6.Scale6(2)
	methods_and_interfaces.ScaleFunc6(&v6, 10)
	p := &methods_and_interfaces.Vertex6{4, 3}
	p.Scale6(3)
	methods_and_interfaces.ScaleFunc6(p, 8)
	fmt.Println("Step 6 -", v6, p)
}


//7. The equivalent thing happens in the reverse direction.
func step7() {
	v7 := methods_and_interfaces.Vertex7{3, 4}
	fmt.Println("Step 7 -", v7.Abs7())
	fmt.Println("Step 7 -", methods_and_interfaces.AbsFunc7(v7))
	//by pointer result similar
	p7 := &methods_and_interfaces.Vertex7{4, 3}
	fmt.Println("Step 7 -", p7.Abs7())
	fmt.Println("Step 7 -", methods_and_interfaces.AbsFunc7(*p7))
}

//8. Here are two reasons to use a pointer receiver.
//The first is so that the method can modify the value that its receiver points to.
//The second is to avoid copying the value on each method call. This can be more efficient
//if the receiver is a large struct, for example.
func step8() {
	fmt.Printf("Step 8\n")
	v := &methods_and_interfaces.Vertex8{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs8())
	v.Scale8(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs8())
}

//9. An interface type is defined as a set of method signatures.
func step9() {
	var a methods_and_interfaces.Abser
	f := methods_and_interfaces.MyFloat9(-math.Sqrt2)
	v := methods_and_interfaces.Vertex9{3, 4}

	a = f  // a MyFloat9 implements Abser
	a = &v // a *Vertex9 implements Abser

	// In the following line, v is a Vertex9 (not *Vertex9)
	// and does NOT implement Abser.
// 		a = v

	fmt.Printf("Step 9\n")
	fmt.Println(a.Abs9())
}

//10. Interfaces are implemented implicitly
func step10() {
	fmt.Printf("Step 10\n")
	var i methods_and_interfaces.I = methods_and_interfaces.T{"hello"}
	i.M()
}

//11. Interface values
func step11() {
	var i methods_and_interfaces.I11

	i = &methods_and_interfaces.T11{"Hello"}
	methods_and_interfaces.Describe11(i)
	i.M11()

	i = methods_and_interfaces.F11(math.Pi)
	methods_and_interfaces.Describe11(i)
	i.M11()
}