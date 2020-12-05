package money

type Money interface {
	Amount() int
	Currency() string
}

func Equals(money, to Money) bool {
	return money.Amount() == to.Amount() && money.Currency() == to.Currency()
}
