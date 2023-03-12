package main

import (
	"os"

	"github.com/pmokeev/CSVCalc/pkg/csv"
)

func main() {
	filepath := os.Args[len(os.Args)-1]

	csvCalculator := csv.NewCSVCalculator()
	csvCalculator.Run(filepath)
}
