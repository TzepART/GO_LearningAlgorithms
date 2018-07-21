package moretypes

import (
	"strings"
	"tutorial/go_packages"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	words := strings.Fields(s)
	for _, w := range words {
		if _, in := m[w]; in {
			m[w] = m[w] + 1
		} else {
			m[w] = 1
		}
	}
	return m
}

func ExerciseMaps23() {
	go_packages.Test(WordCount)
}
