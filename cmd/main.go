package main

import (
	"github.com/alecthomas/kingpin"
	"github.com/pmokeev/CSVCalc/pkg/csv"
)

var (
	filepath = kingpin.Arg("filepath", "Path to CSV file with table.").Required().String()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()

	csvCalculator := csv.NewCSVCalculator()
	csvCalculator.Run(*filepath)
}
