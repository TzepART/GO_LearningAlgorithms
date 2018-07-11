package flowcontrol

import "fmt"

func ForExample1()  {
	sum := 0
	for i := 0; i < 20; i++ {
		sum += i
	}
	fmt.Println(sum)
}
