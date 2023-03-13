package main

import "testing"

func Test_pivotIndex(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			"test1",
			[]int{1, 1, 1, 1, 1, 1},
			-1,
		},
		{
			"test2",
			[]int{1, 7, 3, 6, 5, 6},
			3,
		},
		{
			"test3",
			[]int{1, 2, 3},
			-1,
		},
		{
			"test4",
			[]int{2, 1, -1},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotIndex(tt.nums); got != tt.want {
				t.Errorf("pivotIndex of nums %v = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
