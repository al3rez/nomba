package nomba

type Abser interface {
	Negate() Abser
	LtZero() bool
}

func Abs(x Abser) Abser {
	if x.LtZero() {
		return x.Negate()
	}
	return x
}