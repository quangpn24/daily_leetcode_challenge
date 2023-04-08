package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			"case 1",
			"A man, a plan, a canal: Panama",
			true,
		},
		{
			"case 2",
			"",
			true,
		},
		{
			"case 3",
			"race a car",
			false,
		},
		{
			"case 4",
			"R",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
