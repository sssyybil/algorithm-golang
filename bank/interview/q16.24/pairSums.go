package main

import (
	"sort"
)

//【面试题 16.24. 数对和】 🌀https://leetcode.cn/problems/pairs-with-sum-lcci/description/

func pairSums(nums []int, target int) [][]int {
	var ans [][]int
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	for l < r {
		sum := nums[l] + nums[r]
		if sum == target {
			ans = append(ans, []int{nums[l], nums[r]})
			l++
			r--
		} else if sum < target {
			l++
		} else {
			r--
		}
	}
	return ans
}
