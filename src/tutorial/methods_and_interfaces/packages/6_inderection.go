package methods_and_interfaces

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
