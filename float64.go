package nomba

type Float64 float64

func (f Float64) LtZero() bool {
	return f < 0
}

func (f Float64) Negate() Abser {
	return Float64(f * -1)
}
