package main

import (
	"reflect"
	"testing"
)

func Test_floodFill(t *testing.T) {
	type args struct {
		image [][]int
		sr    int
		sc    int
		color int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"case 1",
			args{
				[][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
				1,
				1,
				2,
			},
			[][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}},
		},
		{
			"case 2",
			args{
				[][]int{{0, 0, 0}, {0, 0, 0}},
				0,
				0,
				0,
			},
			[][]int{{0, 0, 0}, {0, 0, 0}},
		},
		{
			"case 3",
			args{
				[][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
				2,
				2,
				2,
			},
			[][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := floodFill(tt.args.image, tt.args.sr, tt.args.sc, tt.args.color); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("floodFill() = %v, want %v", got, tt.want)
			}
		})
	}
}
