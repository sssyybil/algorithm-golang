package main

import (
	"math"
	"sort"
)

/**
 * 【80. 删除有序数组中的重复项 II】
 * 🪢https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/
 * @2023-04-04 10:33
 */

func removeDuplicates(nums []int) int {
	n := len(nums)
	i, j, res := 0, 0, n
	for i < n-1 {
		counter := 1
		for j < n && nums[i] == nums[j] {
			if counter > 2 {
				nums[j] = math.MaxInt16
			}
			counter++
			j++
		}
		if j-i > 2 {
			res = res - (j - i - 2)
		}
		i = j
	}

	sort.Ints(nums)
	for i, v := range nums {
		if v == math.MaxInt16 {
			nums = nums[:i]
			break
		}
	}
	return res
}
