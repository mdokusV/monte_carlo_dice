package globalVar

import (
	"bufio"
	"os"
)

const NumberOfDices = 2

const MaxNumberOfRounds = 15

const MonteCarloSize = 1_000_000

const DefaultExValue = 3

const MinExValue = 3

const MaxExValue = 6

var Reader *bufio.Reader = bufio.NewReader(os.Stdin)

var Writer *bufio.Writer = bufio.NewWriter(os.Stdout)
