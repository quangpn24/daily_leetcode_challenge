package main

import "testing"

func Test_containsNearbyDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"tc1",
			args{
				nums: []int{1, 2, 3, 1},
				k:    3,
			},
			true,
		},
		{
			"tc2",
			args{
				nums: []int{1, 0, 1, 1},
				k:    1,
			},
			true,
		},
		{
			"tc3",
			args{
				nums: []int{1, 2, 3, 1, 2, 3},
				k:    2,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNearbyDuplicate(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("containsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
