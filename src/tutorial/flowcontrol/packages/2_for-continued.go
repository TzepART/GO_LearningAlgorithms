package flowcontrol

import "fmt"

func ForExample2() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}
