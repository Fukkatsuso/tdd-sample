package money

type Bank struct {
	rates map[Pair]int
}

func NewBank() *Bank {
	return &Bank{
		rates: map[Pair]int{},
	}
}

func (bank *Bank) Reduce(source Expression, to string) *Money {
	return source.Reduce(bank, to)
}

func (bank *Bank) AddRate(from, to string, rate int) {
	bank.rates[NewPair(from, to)] = rate
}

func (bank *Bank) Rate(from, to string) int {
	if from == to {
		return 1
	}
	return bank.rates[NewPair(from, to)]
}
