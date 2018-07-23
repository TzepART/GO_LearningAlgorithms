package methods_and_interfaces

import "fmt"

type I12 interface {
	M12()
}

func NilInterfaceExample12() {
	var i I12
	describe(i)
	i.M12() // panic: runtime error: invalid memory address or nil pointer dereference
}

func describe(i I12) {
	fmt.Printf("(%v, %T)\n", i, i)
}
