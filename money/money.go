package money

type Money struct {
	amount   int
	currency string
}

func NewMoney(amount int, currency string) *Money {
	return &Money{
		amount:   amount,
		currency: currency,
	}
}

func (money *Money) Times(multiplier int) Expression {
	return NewMoney(money.amount*multiplier, money.currency)
}

func (money *Money) Plus(addend Expression) Expression {
	return NewSum(money, addend)
}

func (money *Money) Reduce(bank *Bank, to string) *Money {
	rate := bank.Rate(money.currency, to)
	return NewMoney(money.amount/rate, to)
}

func (money *Money) Currency() string {
	return money.currency
}

func (money *Money) Equals(to *Money) bool {
	return money.amount == to.amount && money.currency == to.currency
}

func NewDollar(amount int) *Money {
	return NewMoney(amount, "USD")
}

func NewFranc(amount int) *Money {
	return NewMoney(amount, "CHF")
}
