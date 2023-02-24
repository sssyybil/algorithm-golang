package main

import "testing"

func Test_titleToNumber(t *testing.T) {
	type args struct {
		columnTitle string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{columnTitle: "A"}, 1},
		{"2", args{columnTitle: "C"}, 3},
		{"3", args{columnTitle: "ABD"}, 732},
		{"4", args{columnTitle: "DD"}, 108},
		{"5", args{columnTitle: "ZY"}, 701},
		{"6", args{columnTitle: "ABC"}, 731},
		{"7", args{columnTitle: "FXSHRXW"}, 2147483647},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := titleToNumber(tt.args.columnTitle); got != tt.want {
				t.Errorf("titleToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
