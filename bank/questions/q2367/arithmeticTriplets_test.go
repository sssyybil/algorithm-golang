package main

import "testing"

func Test_arithmeticTriplets(t *testing.T) {
	type args struct {
		nums []int
		diff int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"01", args{nums: []int{0, 1, 4, 6, 7, 10}, diff: 3}, 2},
		{"02", args{nums: []int{4, 5, 6, 7, 8, 9}, diff: 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arithmeticTriplets(tt.args.nums, tt.args.diff); got != tt.want {
				t.Errorf("arithmeticTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
