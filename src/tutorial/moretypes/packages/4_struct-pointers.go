package moretypes

import "fmt"

type vertex4 struct {
	X int
	Y int
}

func StructExample4() {
	v := vertex4{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
