package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"strings"
	"sync"

	"github.com/mdokusV/monte_carlo_dice/globalVar"
	MonteCarloMath "github.com/mdokusV/monte_carlo_dice/mainMonteCarloMathRand"
	"github.com/mdokusV/monte_carlo_dice/objects"
)

func main() {
	defer globalVar.Writer.Flush()
	f, _ := os.Create("cpu_profile.prof")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	exByRounds := make([][]objects.Round, globalVar.MaxNumberOfRounds)

	for i := range exByRounds {
		exByRounds[i] = make([]objects.Round, globalVar.MaxExValue-globalVar.MinExValue)
		for j := range exByRounds[i] {
			exByRounds[i][j].Number = i

		}
	}

	monteForEx(exByRounds)

	endPrint(exByRounds)

}

func monteForEx(exByRounds [][]objects.Round) {
	var wg sync.WaitGroup
	for exValue := globalVar.MinExValue; exValue < globalVar.MaxExValue; exValue++ {
		wg.Add(1)
		dices, rounds := defineVariables()
		go func(exValue int) {
			defer wg.Done()
			monteCarlo(rounds, dices, exValue, exByRounds)
		}(exValue)
	}
	wg.Wait()
}

func monteCarlo(rounds []objects.Round, dices objects.Dices, exValue int, exByRounds [][]objects.Round) {
	for i := range rounds {
		rounds[i].Zeros()
	}

	MonteCarloMath.InitializeMonteCarlo(dices, rounds, exValue)

	for i := 0; i < globalVar.MaxNumberOfRounds; i++ {
		exByRounds[i][exValue-globalVar.MinExValue] = rounds[i]
	}
}

func endPrint(exByRounds [][]objects.Round) {
	spaces := strings.Repeat(" ", 8)
	fmt.Fprintf(globalVar.Writer, "#  ")

	for i := globalVar.MinExValue; i < globalVar.MaxExValue; i++ {
		fmt.Fprintf(globalVar.Writer, "%d%s", i+1, spaces)
	}
	fmt.Fprintf(globalVar.Writer, "\n")

	for i := range exByRounds {
		fmt.Fprintf(globalVar.Writer, "%d  ", exByRounds[i][0].Number+1)

		for j := range exByRounds[i] {
			exByRounds[i][j].CalculateAverageValueByDiceMonte()
			fmt.Fprintf(globalVar.Writer, "%f ", exByRounds[i][j].Average)
		}
		fmt.Fprintf(globalVar.Writer, "\n")
	}
}

func defineVariables() (objects.Dices, []objects.Round) {
	dices := objects.Dices{
		Dices: make([]int, globalVar.NumberOfDices),
		Sum:   0,
	}

	rounds := make([]objects.Round, globalVar.MaxNumberOfRounds)
	for i := range rounds {
		rounds[i].Number = i
	}
	return dices, rounds
}
