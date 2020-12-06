package money

type Sum struct {
	Augend Expression
	Addend Expression
}

func NewSum(augend Expression, addend Expression) *Sum {
	return &Sum{
		Augend: augend,
		Addend: addend,
	}
}

func (sum *Sum) Times(multiplier int) Expression {
	return NewSum(sum.Augend.Times(multiplier), sum.Addend.Times(multiplier))
}

func (sum *Sum) Plus(addend Expression) Expression {
	return NewSum(sum, addend)
}

func (sum *Sum) Reduce(bank *Bank, to string) *Money {
	amount := sum.Augend.Reduce(bank, to).amount + sum.Addend.Reduce(bank, to).amount
	return NewMoney(amount, to)
}
