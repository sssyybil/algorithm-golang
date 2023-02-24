package main

import "testing"

func Test_numSubarrayProductLessThanK(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"success_01", args{nums: []int{10, 5, 2, 6}, k: 100}, 8},
		{"success_02", args{nums: []int{1, 2, 3}, k: 0}, 0},
		{"success_03", args{nums: []int{1}, k: 10}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubarrayProductLessThanK(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("numSubarrayProductLessThanK_A() = %v, want %v", got, tt.want)
			}
		})
	}
}
