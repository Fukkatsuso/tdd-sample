package money

type Franc Money

func (Franc *Franc) Times(multiplier int) *Money {
	return NewFranc(Franc.amount * multiplier)
}
