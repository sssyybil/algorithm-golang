package main

import "fmt"

/**
 * 【724. 寻找数组的中心下标】🐱https://leetcode.cn/problems/find-pivot-index/
 */

/*
*
==> 设全部元素之和为 total，左侧元素之和为 sum，则右侧元素之和为 total-nums[i]-sum
==> 当 左侧元素之和 == 右侧元素之和时，可以得出 sum == total-nums[i]-sum， 即 total = 2sum + nums[i]
*/
func pivotIndex(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}

	sum := 0
	for i, v := range nums {
		if sum*2+v == total {
			return i
		}
		sum += v
	}
	return -1
}

func main() {
	fmt.Println(
		pivotIndex([]int{1, 7, 3, 6, 5, 6}),
		pivotIndex([]int{1, 2, 3}),
		pivotIndex([]int{2, 1, -1}),
		pivotIndex([]int{1, -1, 2}),
	)
}
