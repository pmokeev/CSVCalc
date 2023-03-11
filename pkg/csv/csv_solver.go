package csv

import (
	"encoding/csv"
	"errors"
	"fmt"
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
	queue            *queue.Queue
	verticalValues   []string
	horisontalValues []string
	cells            map[string][]string
}

func NewCSVCalculator() *CSVCalculator {
	return &CSVCalculator{
		queue:            queue.NewQueue(),
		verticalValues:   make([]string, 0),
		horisontalValues: make([]string, 0),
		cells:            make(map[string][]string, 0),
	}
}

func (cc *CSVCalculator) Run(filepath string) {
	if err := cc.parseCSV(filepath); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	cc.printTable()
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
		return fmt.Errorf("error while reading header line: %v", err.Error())
	}
	cc.horisontalValues = headerRecord

	header, err := cc.createHeaderMap(headerRecord)
	if err != nil {
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

		if len(record)-1 != len(header) {
			return errors.New("size of line doesn't equal header line size")
		}

		values, err := cc.parseLine(record, header)
		if err != nil {
			return err
		}

		cc.cells[record[0]] = values
		cc.verticalValues = append(cc.verticalValues, record[0])
	}

	if err := cc.parseQueue(header); err != nil {
		return err
	}

	return nil
}

func (cc *CSVCalculator) createHeaderMap(record []string) (map[string]int, error) {
	header := make(map[string]int, 0)

	for ind := 1; ind < len(record); ind++ {
		if record[ind] == "" {
			return nil, errors.New("empty header value")
		}

		header[record[ind]] = ind - 1
	}

	return header, nil
}

func (cc *CSVCalculator) parseLine(record []string, header map[string]int) ([]string, error) {
	values := make([]string, 0)

	for ind := 1; ind < len(record); ind++ {
		if strings.ContainsRune(record[ind], '=') {
			term, err := queue.NewTerm(record[ind], record[0], ind-1, header)
			if err != nil {
				return nil, err
			}

			cc.queue.Push(term)
			values = append(values, blank)
		} else {
			if _, err := strconv.Atoi(record[ind]); err != nil {
				return nil, errors.New("unknown value in line")
			}

			values = append(values, record[ind])
		}
	}

	return values, nil
}

func (cc *CSVCalculator) parseQueue(header map[string]int) error {
	for !cc.queue.Empty() {
		term := cc.queue.Pop()

		leftValue, err := term.LeftCell.PickValue(cc.cells)
		if err != nil {
			return err
		}
		if leftValue == blank {
			cc.queue.Push(term)
			continue
		}

		rightValue, err := term.RightCell.PickValue(cc.cells)
		if err != nil {
			return err
		}
		if rightValue == blank {
			cc.queue.Push(term)
			continue
		}

		calculatedValue, err := cc.calculateValue(leftValue, rightValue, term.Operation)
		if err != nil {
			return err
		}

		cc.cells[term.XKey][term.YKey] = calculatedValue
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
			return "", errors.New("division by zero")
		}

		return strconv.Itoa(firstValueConverted / secondValueConverted), nil
	}

	return "", errors.New("strange operation")
}

func (cc *CSVCalculator) printTable() {
	for ind, value := range cc.horisontalValues {
		if ind != len(cc.horisontalValues)-1 {
			fmt.Printf("%v,", value)
		} else {
			fmt.Printf("%v", value)
		}
	}
	fmt.Printf("\n")

	for _, key := range cc.verticalValues {
		fmt.Printf("%v,", key)
		values := cc.cells[key]

		for ind, value := range values {
			if ind != len(values)-1 {
				fmt.Printf("%v,", value)
			} else {
				fmt.Printf("%v", value)
			}
		}

		fmt.Printf("\n")
	}
}
