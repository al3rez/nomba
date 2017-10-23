package nomba

type Float32 float32

func (f Float32) LtZero() bool {
	return f < 0
}

func (f Float32) Negate() Abser {
	return Float32(f * -1)
}
