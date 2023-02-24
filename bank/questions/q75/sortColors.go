package main

import (
	"fmt"
)

/**
 * 【75. 颜色分类】🐱https://leetcode.cn/problems/sort-colors/description/
 */

// 使用【快速排序】
func sortColors(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	p := partition(nums, l, r)
	quickSort(nums, l, p-1)
	quickSort(nums, p+1, r)
}

func partition(nums []int, l, r int) int {
	v, j := nums[l], l
	for i := l + 1; i <= r; i++ {
		if nums[i] < v {
			nums[j+1], nums[i] = nums[i], nums[j+1]
			j++
		}
	}

	nums[l], nums[j] = nums[j], nums[l]
	return j
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}
