package nomba

type Int int

func (i Int) LtZero() bool {
	return i < 0
}

func (i Int) Negate() Abser {
	return Int(i * -1)
}

func (i Int) IsEven() bool {
	return i%2 == 0
}

func (i Int) IsOdd() bool {
	return i%2 != 0
}
