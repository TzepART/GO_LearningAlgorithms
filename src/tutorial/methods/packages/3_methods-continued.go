package methods

type MyFloat float64

func (f MyFloat) Abs3() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
