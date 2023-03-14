package main

import "testing"

func Test_isSubsequence(t *testing.T) {
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
			"test1",
			args{
				s: "add",
				t: "egg",
			},
			false,
		},
		{
			"test2",
			args{
				s: "abc",
				t: "ahbgdc",
			},
			true,
		},
		{
			"test3",
			args{
				s: "axc",
				t: "ahbgdc",
			},
			false,
		},
		{
			"test4",
			args{
				s: "aabb",
				t: "aydabdsb",
			},
			true,
		},
		{
			"test5",
			args{
				s: "aec",
				t: "abcde",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
