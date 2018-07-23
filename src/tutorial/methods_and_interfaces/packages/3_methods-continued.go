package methods_and_interfaces

import (
	"math"
	"fmt"
)

type MyFloat float64

func (f MyFloat) Abs3() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func MethodExample3() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println("Step 3 -", f.Abs3())
}