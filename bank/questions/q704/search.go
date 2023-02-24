package main

import "fmt"

/**
 * 【704. 二分查找】🐱https://leetcode.cn/problems/binary-search/
 */

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(
		search([]int{-1, 0, 3, 5, 9, 12}, 9),
		search([]int{-1, 0, 3, 5, 9, 12}, 2),
	)
}
