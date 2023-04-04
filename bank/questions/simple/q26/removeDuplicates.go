package main

import (
	"math"
	"sort"
)

/**
 * 【26. 删除有序数组中的重复项】
 * 🪢https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/
 * @2023-04-04 14:22
 */

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 1 {
		return n
	}
	i, j, res := 0, 0, n
	for i < n-1 {
		counter := 0
		for j < n && nums[i] == nums[j] {
			if counter >= 1 {
				nums[j] = math.MaxInt16
			}
			counter++
			j++
		}

		if j-i > 1 {
			res = res - (j - i - 1)
		}
		i = j
	}

	sort.Ints(nums)
	nums = nums[:res]
	return res
}
