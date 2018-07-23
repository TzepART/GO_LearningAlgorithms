package methods_and_interfaces

import "fmt"

type I13 interface {
	M12()
}

func NilInterfaceExample13() {
	var i I13
	describe13(i)
	i.M12() // panic: runtime error: invalid memory address or nil pointer dereference
}

func describe13(i I13) {
	fmt.Printf("(%v, %T)\n", i, i)
}
