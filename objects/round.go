package objects

import "github.com/mdokusV/monte_carlo_dice/globalVar"

type Round struct {
	Number  int
	Sum     int
	Average float32
}

func (round *Round) CalculateAverageValueByDiceMonte() {
	round.Average = float32(round.Sum) / float32(globalVar.MonteCarloSize) / float32(globalVar.NumberOfDices)
}

func (round *Round) Zeros() {
	round.Sum = 0
	round.Average = 0
}
