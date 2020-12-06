package money

type Expression interface {
	Times(multiplier int) Expression
	Plus(addend Expression) Expression
	Reduce(bank *Bank, to string) *Money
}
