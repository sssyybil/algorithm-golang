package main

import (
	"algorithm-golang/utils"
	"testing"
)

func Test_runningSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"01", []int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
		{"02", []int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}},
		{"03", []int{3, 1, 2, 10, 1}, []int{3, 4, 6, 16, 17}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runningSum(tt.nums); !utils.IsSliceEquals(got, tt.want) {
				t.Errorf("runningSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
