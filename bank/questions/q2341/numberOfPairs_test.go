package main

import (
	"algorithm-golang/utils"
	"testing"
)

func Test_numberOfPairs(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"success-01", []int{1, 2, 3, 1, 2, 3, 2}, []int{3, 1}},
		{"success-02", []int{1, 2, 3, 1, 2, 3, 2, 2}, []int{4, 0}},
		{"success-03", []int{1, 1}, []int{1, 0}},
		{"success-04", []int{0}, []int{0, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfPairs(tt.nums); !utils.IsSliceEquals(got, tt.want) {
				t.Errorf("numberOfPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
