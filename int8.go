package nomba

type Int8 int8

func (i Int8) LtZero() bool {
	return i < 0
}

func (i Int8) Negate() Abser {
	return Int8(i * -1)
}
