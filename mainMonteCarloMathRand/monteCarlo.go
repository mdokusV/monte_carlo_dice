package MonteCarlo

import (
	"math/rand"
	"sync"
	"time"

	"github.com/mdokusV/monte_carlo_dice/globalVar"
	"github.com/mdokusV/monte_carlo_dice/objects"
)

var randomGenerator *rand.Rand
var randomGeneratorMutex sync.Mutex

func InitializeMonteCarlo(dices objects.Dices, rounds []objects.Round, exValue int) {
	randomGenerator = rand.New(rand.NewSource(time.Now().UnixNano() + int64(exValue)))
	for i := 0; i < globalVar.MonteCarloSize; i++ {
		dices.Zeros()
		throwTurns(&dices, rounds, exValue)
	}
}

func throwTurns(dices *objects.Dices, rounds []objects.Round, exValue int) {
	for round := 0; round < globalVar.MaxNumberOfRounds; round++ {
		change := rollDices(*dices, exValue)
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

func rollDices(dices objects.Dices, expectedValueNatural int) bool {
	change := false
	for i := 0; i < globalVar.NumberOfDices; i++ {
		if dices.Dices[i] <= expectedValueNatural {
			randomGeneratorMutex.Lock()
			dices.Dices[i] = randomGenerator.Intn(6) + 1
			randomGeneratorMutex.Unlock()
			change = true
		}
	}
	return change
}
