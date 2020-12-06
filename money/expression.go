package money

type Expression interface {
	Plus(addend Expression) Expression
	Reduce(bank *Bank, to string) *Money
}
