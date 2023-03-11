package queue

import (
	"reflect"
	"testing"
)

func TestNewCell(t *testing.T) {
	type args struct {
		expression string
		header     map[string]int
	}
	tests := []struct {
		name    string
		args    args
		want    *Cell
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				expression: "Cell30",
				header: map[string]int{
					"Cell": 1,
				},
			},
			want: &Cell{
				XValue: "30",
				YValue: 1,
			},
			wantErr: false,
		},
		{
			name: "Failed: non-existent header value",
			args: args{
				expression: "Cell30",
				header: map[string]int{
					"ABC": 1,
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCell(tt.args.expression, tt.args.header)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCell() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCell() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCell_PickValue(t *testing.T) {
	type fields struct {
		XValue string
		YValue int
	}
	type args struct {
		records map[string][]string
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
				XValue: "30",
				YValue: 0,
			},
			args: args{
				records: map[string][]string{
					"30": {"abc"},
				},
			},
			want:    "abc",
			wantErr: false,
		},
		{
			name: "Failed: non-existent vertical key",
			fields: fields{
				XValue: "30",
				YValue: 1,
			},
			args: args{
				records: make(map[string][]string),
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cell{
				XValue: tt.fields.XValue,
				YValue: tt.fields.YValue,
			}
			got, err := c.PickValue(tt.args.records)
			if (err != nil) != tt.wantErr {
				t.Errorf("Cell.PickValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Cell.PickValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTerm(t *testing.T) {
	type args struct {
		expression string
		xKey       string
		yKey       int
		header     map[string]int
	}
	tests := []struct {
		name    string
		args    args
		want    *Term
		wantErr bool
	}{
		{
			name: "Success: created term",
			args: args{
				expression: "=A1+B2",
				xKey:       "1",
				yKey:       1,
				header: map[string]int{
					"A": 1,
					"B": 2,
				},
			},
			want: &Term{
				XKey: "1",
				YKey: 1,
				LeftCell: &Cell{
					XValue: "1",
					YValue: 1,
				},
				RightCell: &Cell{
					XValue: "2",
					YValue: 2,
				},
				Operation: "+",
			},
			wantErr: false,
		},
		{
			name: "Failed: invalid expression",
			args: args{
				expression: "=A1++B2",
				xKey:       "1",
				yKey:       1,
				header: map[string]int{
					"A": 1,
					"B": 2,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Failed: valid expression but invalid cell",
			args: args{
				expression: "=Cell1+B2",
				xKey:       "1",
				yKey:       1,
				header: map[string]int{
					"A": 1,
					"B": 2,
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTerm(tt.args.expression, tt.args.xKey, tt.args.yKey, tt.args.header)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTerm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTerm() = %v, want %v", got, tt.want)
			}
		})
	}
}
