package money

type Franc struct {
	amount int
}

func NewFranc(amount int) *Franc {
	return &Franc{
		amount: amount,
	}
}

func (franc *Franc) Amount() int {
	return franc.amount
}

func (Franc *Franc) Times(multiplier int) *Franc {
	return NewFranc(Franc.amount * multiplier)
}
