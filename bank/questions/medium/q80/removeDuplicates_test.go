package main

import "testing"

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"01", []int{1, 1, 1, 2, 2, 3}, 5},
		{"02", []int{0, 0, 1, 1, 1, 1, 2, 3, 3}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
