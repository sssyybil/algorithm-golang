package main

import "testing"

func Test_addBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"01", args{a: "11", b: "10"}, "101"},
		{"02", args{a: "11", b: "11"}, "110"},
		{"03", args{a: "1010", b: "1011"}, "10101"},
		{"04", args{a: "1111", b: "11010"}, "101001"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
