package csv

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/pmokeev/CSVCalc/pkg/queue"
)

func TestCSVCalculator_Run(t *testing.T) {
	type fields struct {
		queue *queue.Queue
		table *table
	}
	type args struct {
		filepath string
	}
	tests := []struct {
		name            string
		fields          fields
		args            args
		getCorrectTable func() map[string][]string
		wantErr         bool
	}{
		{
			name: "Success: example.csv",
			fields: fields{
				queue: queue.NewQueue(),
				table: newTable(),
			},
			args: args{
				filepath: "../../test/data/example.csv",
			},
			getCorrectTable: func() map[string][]string {
				return map[string][]string{
					"1":  {"1", "0", "1"},
					"2":  {"2", "6", "0"},
					"30": {"0", "1", "5"},
				}
			},
			wantErr: false,
		},
		{
			name: "Success: one-cell.csv",
			fields: fields{
				queue: queue.NewQueue(),
				table: newTable(),
			},
			args: args{
				filepath: "../../test/data/one-cell.csv",
			},
			getCorrectTable: func() map[string][]string {
				return map[string][]string{
					"1": {"1"},
				}
			},
			wantErr: false,
		},
		{
			name: "Failed: empty-table.csv",
			fields: fields{
				queue: queue.NewQueue(),
				table: newTable(),
			},
			args: args{
				filepath: "../../test/data/empty-table.csv",
			},
			getCorrectTable: func() map[string][]string {
				return make(map[string][]string, 0)
			},
			wantErr: true,
		},
		{
			name: "Failed: division-by-zero.csv",
			fields: fields{
				queue: queue.NewQueue(),
				table: newTable(),
			},
			args: args{
				filepath: "../../test/data/division-by-zero.csv",
			},
			getCorrectTable: func() map[string][]string {
				return map[string][]string{
					"1":  {"0", "0", "1"},
					"2":  {"2", "5", "0"},
					"30": {"0", "_", "5"},
				}
			},
			wantErr: true,
		},
		{
			name: "Failed: empty-header-value.csv",
			fields: fields{
				queue: queue.NewQueue(),
				table: newTable(),
			},
			args: args{
				filepath: "../../test/data/empty-header-value.csv",
			},
			getCorrectTable: func() map[string][]string {
				return make(map[string][]string, 0)
			},
			wantErr: true,
		},
		{
			name: "Failed: different-lines-length.csv",
			fields: fields{
				queue: queue.NewQueue(),
				table: newTable(),
			},
			args: args{
				filepath: "../../test/data/different-lines-length.csv",
			},
			getCorrectTable: func() map[string][]string {
				return make(map[string][]string, 0)
			},
			wantErr: true,
		},
		{
			name: "Failed: cyclic-dependency.csv",
			fields: fields{
				queue: queue.NewQueue(),
				table: newTable(),
			},
			args: args{
				filepath: "../../test/data/cyclic-dependency.csv",
			},
			getCorrectTable: func() map[string][]string {
				return map[string][]string{
					"1": {"1", "_"},
					"2": {"_", "3"},
				}
			},
			wantErr: true,
		},
		{
			name: "Failed: incorrect-cell-reference.csv",
			fields: fields{
				queue: queue.NewQueue(),
				table: newTable(),
			},
			args: args{
				filepath: "../../test/data/incorrect-cell-reference.csv",
			},
			getCorrectTable: func() map[string][]string {
				return map[string][]string{
					"1": {"1", "0", "1"},
					"2": {"2", "_", "0"},
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := &CSVCalculator{
				queue: tt.fields.queue,
				table: tt.fields.table,
			}
			err := cc.Run(tt.args.filepath)

			fmt.Printf("%v", cc.table.cells)

			if (err != nil) != tt.wantErr {
				t.Errorf("CSVCalculator.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(tt.getCorrectTable(), cc.table.cells) {
				t.Errorf("CSVCalculator.Run() unequal tables")
			}
		})
	}
}

func TestCSVCalculator_parseLine(t *testing.T) {
	type fields struct {
		queue *queue.Queue
		table *table
	}
	type args struct {
		record []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				queue: queue.NewQueue(),
				table: newTable(),
			},
			args: args{
				record: []string{"1", "2", "3"},
			},
			want:    []string{"2", "3"},
			wantErr: false,
		},
		{
			name: "Success: with blank value",
			fields: fields{
				queue: queue.NewQueue(),
				table: &table{
					header: map[string]int{
						"A": 1,
					},
				},
			},
			args: args{
				record: []string{"1", "=A1+A2", "3"},
			},
			want:    []string{blank, "3"},
			wantErr: false,
		},
		{
			name: "Failed: invalid expression",
			fields: fields{
				queue: queue.NewQueue(),
				table: newTable(),
			},
			args: args{
				record: []string{"1", "=++", "3"},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := &CSVCalculator{
				queue: tt.fields.queue,
				table: tt.fields.table,
			}
			got, err := cc.parseLine(tt.args.record)
			if (err != nil) != tt.wantErr {
				t.Errorf("CSVCalculator.parseLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CSVCalculator.parseLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCSVCalculator_calculateValue(t *testing.T) {
	type fields struct {
		queue *queue.Queue
		table *table
	}
	type args struct {
		firstValue  string
		secondValue string
		operation   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				queue: queue.NewQueue(),
				table: newTable(),
			},
			args: args{
				firstValue:  "1",
				secondValue: "1",
				operation:   "+",
			},
			want:    "2",
			wantErr: false,
		},
		{
			name: "Failed: unknown operation",
			fields: fields{
				queue: queue.NewQueue(),
				table: newTable(),
			},
			args: args{
				firstValue:  "1",
				secondValue: "1",
				operation:   "123",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Failed: division by zero",
			fields: fields{
				queue: queue.NewQueue(),
				table: newTable(),
			},
			args: args{
				firstValue:  "1",
				secondValue: "0",
				operation:   "/",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := &CSVCalculator{
				queue: tt.fields.queue,
				table: tt.fields.table,
			}
			got, err := cc.calculateValue(tt.args.firstValue, tt.args.secondValue, tt.args.operation)
			if (err != nil) != tt.wantErr {
				t.Errorf("CSVCalculator.calculateValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CSVCalculator.calculateValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
