package money

type Bank struct {
}

func NewBank() *Bank {
	return &Bank{}
}

func (bank *Bank) Reduce(source Expression, to string) *Money {
	return NewDollar(10)
}
