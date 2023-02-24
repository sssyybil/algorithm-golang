package helper

import (
	"fmt"
	"math/rand"
)

// GenerateRandomArray 可生成范围在 [rangeL,rangeR] 区间、包含 n 个元素的数组
func GenerateRandomArray(n, rangeL, rangeR int) (nums []int) {
	if rangeL >= rangeR {
		fmt.Println(fmt.Errorf("数组范围有误，rangeL 必须大于 rangeR，rangeL=%d, rangeR=%d", rangeL, rangeR))
		return
	}
	nums = make([]int, n)
	for i := 0; i < n; i++ {
		v := int(rand.Float64()*float64(rangeR-rangeL+1)) + rangeL
		nums[i] = v
	}
	return
}

// GenerateNearlyOrderArray 可生成长度为 n 的近乎有序的数组
func GenerateNearlyOrderArray(n, swapTimes int) (nums []int) {
	nums = make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i
	}

	for i := 0; i < swapTimes; i++ {
		indexA := int(rand.Float64() * float64(n))
		indexB := int(rand.Float64() * float64(n))

		nums[indexA], nums[indexB] = nums[indexB], nums[indexA]
	}

	return nums
}

// IsSorted 判断数组是否有序
// true: 有序；false：无序
func IsSorted(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return false
		}
	}
	return true
}
