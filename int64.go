package nomba

type Int64 int64

func (i Int64) LtZero() bool {
	return i < 0
}

func (i Int64) Negate() Abser {
	return Int64(i * -1)
}

func (i Int64) IsEven() bool {
	return i%2 == 0
}

func (i Int64) IsOdd() bool {
	return i%2 != 0
}
