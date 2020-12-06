package money

type Dollar Money

func (dollar *Dollar) Times(multiplier int) *Money {
	return NewDollar(dollar.amount * multiplier)
}
