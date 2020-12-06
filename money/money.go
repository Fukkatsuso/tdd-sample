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
