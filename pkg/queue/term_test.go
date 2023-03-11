package queue

import (
	"reflect"
	"testing"
)

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
					YValue: "A",
				},
				RightCell: &Cell{
					XValue: "2",
					YValue: "B",
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
