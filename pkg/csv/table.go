package csv

import "fmt"

type table struct {
	header         map[string]int
	horizontalKeys []string
	verticalKeys   []string
	cells          map[string][]string
}

func newTable() *table {
	return &table{
		verticalKeys:   make([]string, 0),
		horizontalKeys: make([]string, 0),
		header:         make(map[string]int, 0),
		cells:          make(map[string][]string, 0),
	}
}

func (t *table) print() {
	for ind, value := range t.horizontalKeys {
		if ind != len(t.horizontalKeys)-1 {
			fmt.Printf("%v,", value)
		} else {
			fmt.Printf("%v", value)
		}
	}
	fmt.Printf("\n")

	for _, key := range t.verticalKeys {
		fmt.Printf("%v,", key)
		values := t.cells[key]

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

func (t *table) setHeader(value []string) error {
	t.horizontalKeys = value

	header := make(map[string]int, 0)

	for ind := 1; ind < len(t.horizontalKeys); ind++ {
		if t.horizontalKeys[ind] == "" {
			return errEmptyValueInHeader
		}

		header[t.horizontalKeys[ind]] = ind - 1
	}

	t.header = header

	return nil
}

func (t *table) addHorizontalLine(key string, values []string) {
	t.cells[key] = values
	t.verticalKeys = append(t.verticalKeys, key)
}

func (t *table) setCellValue(xKey int, yKey, value string) {
	t.cells[yKey][xKey] = value
}
