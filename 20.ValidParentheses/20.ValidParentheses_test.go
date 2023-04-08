package main

import "testing"

func Test_isValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			"case 1",
			"{}(())",
			true,
		},
		{
			"case 2",
			"{}()[]",
			true,
		},
		{
			"case 3",
			"{]",
			false,
		},
		{
			"case 4",
			"({[}])",
			false,
		},
		{
			"case 5",
			"{[]}",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
