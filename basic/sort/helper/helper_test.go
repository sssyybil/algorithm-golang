package helper

import (
	"fmt"
	"testing"
)

func TestGenerateRandomArray(t *testing.T) {
	nums := GenerateRandomArray(10, 5, 15)
	fmt.Println(
		nums,
		IsSorted(nums),
	)

}

func TestGenerateNearlyOrderArray(t *testing.T) {
	fmt.Println(
		GenerateNearlyOrderArray(8, 3),
		GenerateNearlyOrderArray(6, 0),
	)
}
