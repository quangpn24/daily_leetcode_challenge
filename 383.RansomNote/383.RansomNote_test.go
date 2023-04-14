package main

import "testing"

func Test_canConstruct(t *testing.T) {
	type args struct {
		ransomNote string
		magazine   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"case 1",
			args{
				ransomNote: "aa",
				magazine:   "ab",
			},
			false,
		},
		{
			"case 2",
			args{
				ransomNote: "aa",
				magazine:   "aab",
			},
			true,
		},
		{
			"case 3",
			args{
				ransomNote: "a",
				magazine:   "b",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstruct(tt.args.ransomNote, tt.args.magazine); got != tt.want {
				t.Errorf("canConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
