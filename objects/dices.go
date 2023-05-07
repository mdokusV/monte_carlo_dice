package objects

type Dices struct {
	Dices []int
	Sum   int
}

func (dices *Dices) CalculateSum() {
	dices.Sum = 0
	for _, dice := range dices.Dices {
		dices.Sum += dice
	}
}

func (dices *Dices) Zeros() {
	for i := range dices.Dices {
		dices.Dices[i] = 0
	}
}
