package moretypes

import "fmt"

type myVertex2 struct {
	N int
	M int
}

type vertex2 struct {
	X int
	Y int
	z string
	MV myVertex2
}

func StructExample2() {
	mv := myVertex2{4,5}
	fmt.Println(vertex2{1, 2, "Hello!", mv})
}
