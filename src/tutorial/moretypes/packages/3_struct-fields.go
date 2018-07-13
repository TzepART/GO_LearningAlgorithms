package moretypes

import "fmt"

type vertex3 struct {
	X int
	Y int
}

func StructExample3() {
	v := vertex3{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
