package money

type Money interface {
	Amount() int
}

func Equals(money, to Money) bool {
	return money.Amount() == to.Amount()
}
