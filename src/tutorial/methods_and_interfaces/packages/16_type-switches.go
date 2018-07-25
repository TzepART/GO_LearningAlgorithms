package methods_and_interfaces

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	case float64:
		fmt.Printf("Add 10 to %v is %v\n", v, v+10)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func TypeSwitchesExample16() {
	do(21)
	do("hello")
	do(7.63)
	do(true)
}
