package main

import "testing"

func Test_isAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"case1",
			args{
				s: "anagram",
				t: "nagaram",
			},
			true,
		},
		{
			"case2",
			args{
				s: "anagram",
				t: "nagar",
			},
			false,
		},
		{
			"case3",
			args{
				s: "cat",
				t: "rat",
			},
			false,
		},
		{
			"case4",
			args{
				s: "abcs",
				t: "ssab",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
