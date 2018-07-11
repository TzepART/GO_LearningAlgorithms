package methods

import (
	"fmt"
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
