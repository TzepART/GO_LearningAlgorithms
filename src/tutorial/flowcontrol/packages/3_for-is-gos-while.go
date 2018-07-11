package flowcontrol

import "fmt"

func ForExample3() {
	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)
}
