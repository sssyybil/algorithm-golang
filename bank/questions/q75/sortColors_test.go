package main

import (
	"algorithm-golang/basic/sort/helper"
	"testing"
)

func Test_sortColors(t *testing.T) {

	arr1 := []int{2, 0, 2, 1, 1, 0}
	arr2 := []int{2, 0, 1}
	arr3 := []int{2, 2, 1, 1, 0, 0, 0}
	arr4 := []int{0, 1, 2, 0, 1, 2}

	tests := []struct {
		name string
		nums []int
	}{
		{"01", arr1},
		{"02", arr2},
		{"03", arr3},
		{"04", arr4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.nums)
			if !helper.IsSorted(tt.nums) {
				t.Errorf("not sorted. nums=%v\n", tt.nums)
			}

		})
	}
}

func Test_sortColorsWithThreeWayQs(t *testing.T) {
	arr1 := []int{2, 0, 2, 1, 1, 0}
	arr2 := []int{2, 0, 1}
	arr3 := []int{2, 2, 1, 1, 0, 0, 0}
	arr4 := []int{0, 1, 2, 0, 1, 2}

	tests := []struct {
		name string
		nums []int
	}{
		{"01", arr1},
		{"02", arr2},
		{"03", arr3},
		{"04", arr4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColorsWithThreeWayQs(tt.nums)
			if !helper.IsSorted(tt.nums) {
				t.Errorf("not sorted. nums=%v\n", tt.nums)
			}

		})
	}
}
