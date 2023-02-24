package main

import "testing"

func Test_divide(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"01", args{a: 15, b: 2}, 7},
		{"02", args{a: 7, b: -3}, -2},
		{"03", args{a: 0, b: 1}, 0},
		{"04", args{a: 1, b: 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
