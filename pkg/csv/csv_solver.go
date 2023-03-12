package csv

import (
	"encoding/csv"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/pmokeev/CSVCalc/pkg/queue"
)

const (
	epsilon = 1e-10
	blank   = "_"
)

type CSVCalculator struct {
	queue *queue.Queue
	table *table
}

func NewCSVCalculator() *CSVCalculator {
	return &CSVCalculator{
		queue: queue.NewQueue(),
		table: newTable(),
	}
}

func (cc *CSVCalculator) Run(filepath string) {
	if err := cc.parseCSV(filepath); err != nil {
		log.Fatal(err)
	}

	if err := cc.parseQueue(); err != nil {
		log.Fatal(err)
	}

	cc.table.print()
}

func (cc *CSVCalculator) parseCSV(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)

	headerRecord, err := csvReader.Read()
	if err != nil {
		return err
	}

	if err := cc.table.setHeader(headerRecord); err != nil {
		return err
	}

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if len(record)-1 != len(cc.table.header) {
			return errUnequalLinesLengths
		}

		values, err := cc.parseLine(record)
		if err != nil {
			return err
		}

		cc.table.addHorizontalLine(record[0], values)
	}

	return nil
}

func (cc *CSVCalculator) parseLine(record []string) ([]string, error) {
	values := make([]string, 0)

	for ind := 1; ind < len(record); ind++ {
		if strings.ContainsRune(record[ind], '=') {
			term, err := queue.NewTerm(record[ind], record[0], ind-1, cc.table.header)
			if err != nil {
				return nil, err
			}

			cc.queue.Push(term)
			values = append(values, blank)
		} else {
			if _, err := strconv.Atoi(record[ind]); err != nil {
				return nil, &errUnknownValueInLine{value: record[ind]}
			}

			values = append(values, record[ind])
		}
	}

	return values, nil
}

func (cc *CSVCalculator) parseQueue() error {
	waitingCells := make(map[string]bool, 0)

	for !cc.queue.Empty() {
		term := cc.queue.Pop()

		leftValue, err := term.LeftCell.PickValue(cc.table.cells)
		if err != nil {
			return err
		}
		if leftValue == blank {
			cc.queue.Push(term)
			waitingCells[term.LeftCell.String()] = true
			continue
		}

		rightValue, err := term.RightCell.PickValue(cc.table.cells)
		if err != nil {
			return err
		}
		if rightValue == blank {
			cc.queue.Push(term)
			waitingCells[term.RightCell.String()] = true
			continue
		}

		calculatedValue, err := cc.calculateValue(leftValue, rightValue, term.Operation)
		if err != nil {
			return err
		}

		currentCell := &queue.Cell{
			XValue: term.XKey,
			YValue: term.YKey,
		}

		if _, ok := waitingCells[currentCell.String()]; ok {
			return errCyclicDependency
		}

		cc.table.setCellValue(term.XKey, term.YKey, calculatedValue)
		delete(waitingCells, currentCell.String())
	}

	return nil
}

func (cc *CSVCalculator) calculateValue(firstValue, secondValue, operation string) (string, error) {
	firstValueConverted, err := strconv.Atoi(firstValue)
	if err != nil {
		return "", err
	}

	secondValueConverted, err := strconv.Atoi(secondValue)
	if err != nil {
		return "", err
	}

	switch operation {
	case "+":
		return strconv.Itoa(firstValueConverted + secondValueConverted), nil
	case "-":
		return strconv.Itoa(firstValueConverted - secondValueConverted), nil
	case "*":
		return strconv.Itoa(firstValueConverted * secondValueConverted), nil
	case "/":
		if math.Abs(float64(secondValueConverted)) < epsilon {
			return "", errDivisionByZero
		}

		return strconv.Itoa(firstValueConverted / secondValueConverted), nil
	}

	return "", &queue.ErrUnknownOperationInExpression{Value: operation}
}
