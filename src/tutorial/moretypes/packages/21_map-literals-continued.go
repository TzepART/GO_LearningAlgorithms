package moretypes

import "fmt"

type Vertex21 struct {
	Lat, Long float64
}

func MapExample21() {

	var m = map[string]Vertex21{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	fmt.Println(m)
}
