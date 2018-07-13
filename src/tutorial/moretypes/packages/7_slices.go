package moretypes

import "fmt"

func SliceExample7() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s = primes[2:4]
	fmt.Println(s)
}
