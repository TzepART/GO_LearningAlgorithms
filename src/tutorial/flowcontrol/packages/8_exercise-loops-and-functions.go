package flowcontrol

import (
	"fmt"
)

func Sqrt8(x float64) (z float64) {
	z=x
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}

	return
}

func Exercise8() {
	fmt.Println(Sqrt8(2))
}
