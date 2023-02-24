package quicksort

import (
	"algorithm-golang/basic/sort/helper"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr1 := helper.GenerateRandomArray(10, 1, 99)
	arr2 := helper.GenerateRandomArray(1000, 3, 999)
	arr3 := helper.GenerateNearlyOrderArray(100, 3)
	arr4 := helper.GenerateNearlyOrderArray(1000, 0)

	tests := []struct {
		name string
		arr  []int
	}{
		{"randomArray-01", arr1},
		{"randomArray-02", arr2},
		{"nearlyOrderArray-01", arr3},
		{"nearlyOrderArray-02", arr4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sort(tt.arr)
			if !helper.IsSorted(tt.arr) {
				t.Errorf("QuickSorted res=%v\n", tt.arr)
			}
		})
	}
}

func TestDualPivotQuickSort(t *testing.T) {
	arr1 := helper.GenerateRandomArray(8, 1, 20)
	arr2 := helper.GenerateRandomArray(1000, 3, 999)
	arr3 := helper.GenerateNearlyOrderArray(100, 3)
	arr4 := helper.GenerateNearlyOrderArray(1000, 0)

	tests := []struct {
		name string
		arr  []int
	}{
		{"randomArray-01", arr1},
		{"randomArray-02", arr2},
		{"nearlyOrderArray-01", arr3},
		{"nearlyOrderArray-02", arr4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DualPivotQuickSort(tt.arr)
			if !helper.IsSorted(tt.arr) {
				t.Errorf("QuickSorted res=%v\n", tt.arr)
			}
		})
	}
}
