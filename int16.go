package nomba

type Int16 int16

func (i Int16) LtZero() bool {
	return i < 0
}

func (i Int16) Negate() Abser {
	return Int16(i * -1)
}
