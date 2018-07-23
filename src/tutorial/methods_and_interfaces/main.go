package main

import (
	"tutorial/methods_and_interfaces/packages"
)

// https://tour.golang.org/methods_and_interfaces/1
func main() {
	step12()
}

//1. Use method is a function with a special receiver argument.
func step1()  {
	methods_and_interfaces.MethodExample1()
}

//2. Use method is a function without a special receiver argument.
func step2() {
	methods_and_interfaces.MethodExample2()
}

//3. Declare a method on non-struct types
func step3() {
	methods_and_interfaces.MethodExample3()
}

//4. Declare methods_and_interfaces with pointer receivers
func step4() {
	methods_and_interfaces.MethodPointer4()
}

//5. Here we see the Abs and Scale methods_and_interfaces rewritten as functions
func step5() {
	methods_and_interfaces.MethodPointer5()
}

//6. methods_and_interfaces and pointer indirection
func step6() {
	methods_and_interfaces.MethodInderection6()
}

//7. The equivalent thing happens in the reverse direction.
func step7() {
	methods_and_interfaces.MethodInderection7()
}

//8. Choosing a value or pointer receiver
func step8() {
	methods_and_interfaces.MethodPointer8()
}

//9. An interface type is defined as a set of method signatures.
func step9() {
	methods_and_interfaces.InterfaceExample9()
}

//10. Interfaces are implemented implicitly
func step10() {
	methods_and_interfaces.InterfaceExample10()
}

//11. Interface values
func step11() {
	methods_and_interfaces.InterfaceValueExample11()
}

//12. Nil interface values
func step12()  {
	methods_and_interfaces.NilInterfaceExample12()
}

////13.
//func step13()  {
//	methods_and_interfaces.Example13()
//}
//
////14.
//func step14()  {
//	methods_and_interfaces.Example14()
//}
//
////15.
//func step15()  {
//	methods_and_interfaces.Example15()
//}
//
////16.
//func step16()  {
//	methods_and_interfaces.Example16()
//}
//
////17.
//func step17()  {
//	methods_and_interfaces.Example17()
//}
//
////18.
//func step18()  {
//	methods_and_interfaces.Example18()
//}
//
////19.
//func step19()  {
//	methods_and_interfaces.Example19()
//}
//
////20.
//func step20()  {
//	methods_and_interfaces.Example20()
//}
//
////21.
//func step21()  {
//	methods_and_interfaces.Example21()
//}
//
////22.
//func step22()  {
//	methods_and_interfaces.Example22()
//}
//
////23.
//func step23()  {
//	methods_and_interfaces.Example23()
//}
//
////24.
//func step24()  {
//	methods_and_interfaces.Example24()
//}
//
////25.
//func step25()  {
//	methods_and_interfaces.Example25()
//}