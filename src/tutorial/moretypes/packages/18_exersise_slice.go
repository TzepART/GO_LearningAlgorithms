package moretypes

import "tutorial/go_packages"

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)

	for x := range pic {
		pic[x] = make([]uint8, dx)
		for y := range pic[x] {
			pic[x][y] = (uint8(y) << (uint8(x)/32))
		}
	}
	return pic
}

func ExerciseSlice18() {
	go_packages.ShowPic(Pic)
}
