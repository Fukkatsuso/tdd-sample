package money

type Money interface {
	Amount() int
	Currency() string
	Times(multiplier int) Money
}

func Equals(money, to Money) bool {
	return money.Amount() == to.Amount() && money.Currency() == to.Currency()
}
