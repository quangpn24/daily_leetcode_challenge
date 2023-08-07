package main

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{
				nums:   []int{1, 2, 3, 4, 5},
				target: 5,
			},
			4,
		},
		{
			"test2",
			args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 9,
			},
			4,
		},
		{
			"test3",
			args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 2,
			},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}