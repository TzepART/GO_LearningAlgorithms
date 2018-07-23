package methods_and_interfaces

import (
	"fmt"
	"math"
)

type I11 interface {
	M11()
}

type T11 struct {
	S string
}

func (t *T11) M11() {
	fmt.Println(t.S)
}

type F11 float64

func (f F11) M11() {
	fmt.Println(f)
}

func Describe11(i I11) {
	fmt.Printf("(%v, %T)\n", i, i)
}

//11. Interface values
func InterfaceValueExample11() {
	var i I11

	i = &T11{"Hello"}
	Describe11(i)
	i.M11()

	i = F11(math.Pi)
	Describe11(i)
	i.M11()
}
