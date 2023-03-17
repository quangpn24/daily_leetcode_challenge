package main

import "testing"

func Test_maxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			"test1",
			[]int{7, 1, 5, 3, 6, 4},
			5,
		},
		{
			"test2",
			[]int{1, 1, 2, 5, 3, 7, 4},
			6,
		},
		{
			"test3",
			[]int{1, 1, 1, 1},
			0,
		},
		{
			"test4",
			[]int{2, 4, 1},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
