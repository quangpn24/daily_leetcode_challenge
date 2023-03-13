package main

import (
	"reflect"
	"testing"
)

func Test_runningSum(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			"test1",
			[]int{1, 2, 3, 4},
			[]int{1, 3, 6, 10},
		},
		{
			"test2",
			[]int{1, 1, 1, 1, 1},
			[]int{1, 2, 3, 4, 5},
		},
		{
			"test3",
			[]int{-1, 2, -4, 2, 1},
			[]int{-1, 1, -3, -1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runningSum(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("runningSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
