package main

import "testing"

func Test_maxArea(t *testing.T) {

	tests := []struct {
		name   string
		height []int
		want   int
	}{
		// TODO: Add test cases.
		{
			"test1",
			[]int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			49,
		},
		{
			"test2",
			[]int{1, 1},
			1,
		},
		{
			"test3",
			[]int{1, 0, 1},
			2,
		},
		{
			"test4",
			[]int{0, 1, 0},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
