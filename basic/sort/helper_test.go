package main

import (
	"algorithm-golang/basic/sort/helper"
	"fmt"
	"testing"
)

func TestGenerateRandomArray(t *testing.T) {
	nums := helper.GenerateRandomArray(10, 5, 15)
	fmt.Println(
		nums,
		helper.IsSorted(nums),
	)

}

func TestGenerateNearlyOrderArray(t *testing.T) {
	fmt.Println(
		helper.GenerateNearlyOrderArray(8, 3),
		helper.GenerateNearlyOrderArray(6, 0),
	)
}
