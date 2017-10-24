package nomba

type Int8 int8

func (i Int8) LtZero() bool {
	return i < 0
}

func (i Int8) Negate() Abser {
	return Int8(i * -1)
}

func (i Int8) IsEven() bool {
	return i%2 == 0
}

func (i Int8) IsOdd() bool {
	return i%2 != 0
}
