package main

import (
	"log"
	"os"

	"github.com/pmokeev/CSVCalc/pkg/csv"
)

func main() {
	filepath := os.Args[len(os.Args)-1]

	csvCalculator := csv.NewCSVCalculator()
	if err := csvCalculator.Run(filepath); err != nil {
		log.Fatal(err)
	}
}
