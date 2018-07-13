package moretypes

import "fmt"

type MyVertex struct {
	N int
	M int
}

type Vertex struct {
	X int
	Y int
	z string
	MV MyVertex
}

func StructExample2() {
	mv := MyVertex{4,5}
	fmt.Println(Vertex{1, 2, "Hello!", mv})
}
