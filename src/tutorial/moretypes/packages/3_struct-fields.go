package moretypes

import "fmt"

type vertex struct {
	X int
	Y int
}

func StructExample3() {
	v := vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
