package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		args int
		want bool
	}{
		{
			"test1",
			121,
			true,
		},
		{
			"test2",
			-121,
			false,
		},
		{
			"test3",
			12345677654321,
			true,
		},
		{
			"test4",
			11111111,
			true,
		},
		{
			"test5",
			123113,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
