package main

import "testing"

func Test_isIsomorphic(t *testing.T) {
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
			true,
		},
		{
			"test2",
			args{
				s: "foo",
				t: "bar",
			},
			false,
		},
		{
			"test3",
			args{
				s: "paper",
				t: "title",
			},
			true,
		},
		{
			"test4",
			args{
				s: "aaabbbba",
				t: "bbbaaaba",
			},
			false,
		},
		{
			"test5",
			args{
				s: "abababab",
				t: "bcbcbcbc",
			},
			true,
		},
		{
			"test6",
			args{
				s: "badc",
				t: "baba",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphic(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}
