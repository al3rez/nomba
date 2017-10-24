package nomba

type Int32 int32

func (i Int32) LtZero() bool {
	return i < 0
}

func (i Int32) Negate() Abser {
	return Int32(i * -1)
}

func (i Int32) IsEven() bool {
	return i%2 == 0
}

func (i Int32) IsOdd() bool {
	return i%2 != 0
}
