package methods_and_interfaces

import "fmt"

type Vertex6 struct {
	X, Y float64
}

func (v *Vertex6) Scale6(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc6(v *Vertex6, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func MethodInderection6() {
	v6 := Vertex6{3, 4}
	v6.Scale6(2)
	ScaleFunc6(&v6, 10)
	p := &Vertex6{4, 3}
	p.Scale6(3)
	ScaleFunc6(p, 8)
	fmt.Println("Step 6 -", v6, p)
}
