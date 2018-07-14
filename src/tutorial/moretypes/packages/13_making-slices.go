package moretypes

import "fmt"

func SliceExample13() {
	a := make([]int, 5)
	printSlice13("a", a)

	b := make([]int, 0, 5)
	printSlice13("b", b)

	c := b[:2]
	printSlice13("c", c)

	d := c[2:5]
	printSlice13("d", d)
}

func printSlice13(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
