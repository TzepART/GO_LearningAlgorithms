package moretypes

import "fmt"

type vertex5 struct {
	X, Y int
}

var (
	v1 = vertex5{1, 2}  // has type vertex5
	v2 = vertex5{X: 1}  // Y:0 is implicit
	v3 = vertex5{}      // X:0 and Y:0
	p  = &vertex5{1, 2} // has type *vertex5
)

func StructExample5() {
	fmt.Println(v1, p, v2, v3)
}
