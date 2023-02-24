package insertsort

import (
	"algorithm-golang/basic/sort/helper"
	"testing"
)

func TestSortWithCopy(t *testing.T) {
	tests := []struct {
		name string
		nums []int
	}{
		// TODO: Add test cases.
		{"1", helper.GenerateRandomArray(10, 1, 99)},
		{"2", helper.GenerateRandomArray(100, 9, 999)},
		{"3", helper.GenerateRandomArray(1000, 1, 9999)},
		{"4", helper.GenerateRandomArray(10000, 1, 998)},
		{"5", helper.GenerateRandomArray(100000, 1, 9988)},
		{"6", helper.GenerateRandomArray(1000000, 1, 9900)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortWithCopySimpler(tt.nums)
			if !helper.IsSorted(tt.nums) {
				t.Errorf("sorting fail, result: %v", tt.nums)
			}
		})
	}
}
