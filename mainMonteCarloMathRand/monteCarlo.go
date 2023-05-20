package MonteCarlo

import (
	"math/rand"
	"time"

	ini "github.com/mdokusV/monte_carlo_dice/Ini"
	"github.com/mdokusV/monte_carlo_dice/globalVar"
	"github.com/mdokusV/monte_carlo_dice/objects"
)

func InitializeMonteCarlo(dices objects.Dices, rounds []objects.Round, exValue int) {
	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano() + int64(exValue)))
	for i := 0; i < globalVar.MonteCarloSize; i++ {
		dices.Zeros()
		throwTurns(&dices, rounds, exValue, randomGenerator)
	}
}

func throwTurns(dices *objects.Dices, rounds []objects.Round, exValue int, randomGenerator *rand.Rand) {
	for round := 0; round < globalVar.MaxNumberOfRounds; round++ {
		change := rollDices(*dices, exValue, randomGenerator)
		dices.CalculateSum()
		rounds[round].Sum += dices.Sum
		if !change {
			for easyRound := round + 1; easyRound < globalVar.MaxNumberOfRounds; easyRound++ {
				rounds[easyRound].Sum += dices.Sum
			}
			break
		}
	}
}

func MonteUsingCalcEx(exByRounds [][]objects.Round) {
	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	dices, rounds := ini.DefineVariables()
	for i := globalVar.MaxNumberOfRounds - 1; i >= 0; i-- {
		change := rollDices(*dices, exValue, randomGenerator)

	}

}

func rollDices(dices objects.Dices, expectedValueNatural int, randomGenerator *rand.Rand) bool {
	change := false
	for i := 0; i < globalVar.NumberOfDices; i++ {
		if dices.Dices[i] <= expectedValueNatural {
			dices.Dices[i] = randomGenerator.Intn(6) + 1
			change = true
		}
	}
	return change
}
