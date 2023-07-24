package main

import (
	"math"
	"testing"
)

func Test_reverse(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{
			"tc_happy_case",
			123,
			321,
		},
		{
			"tc1_happy_case",
			1234567,
			7654321,
		},
		{
			"tc_negative",
			-123,
			-321,
		},
		{
			"> 2^31 - 1",
			int(math.Pow(2, 31)),
			0,
		},
		{
			"< -2^31",
			int(-math.Pow(2, 31) - 1),
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
