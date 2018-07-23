package methods_and_interfaces

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func InterfaceExample10() {
	fmt.Printf("Step 10\n")
	var i I = T{"hello"}
	i.M()
}