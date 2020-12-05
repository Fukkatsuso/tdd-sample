package money

type Dollar struct {
	amount int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{
		amount: amount,
	}
}

func (dollar *Dollar) Amount() int {
	return dollar.amount
}

func (dollar *Dollar) Currency() string {
	return "dollar"
}

func (dollar *Dollar) Times(multiplier int) Money {
	return NewDollar(dollar.amount * multiplier)
}
