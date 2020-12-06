package money

type Franc Money

func (franc *Franc) Times(multiplier int) *Money {
	return NewFranc(franc.amount * multiplier)
}
