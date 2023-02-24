package main

import "fmt"

/**
 * 【1480. 一维数组的动态和】🐱https://leetcode.cn/problems/running-sum-of-1d-array/
 */

func runningSum(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	res[0] = nums[0]
	if n == 1 {
		return res
	}
	curSum := nums[0]
	for i := 1; i < n; i++ {
		res[i] = curSum + nums[i]
		curSum = res[i]
	}
	return res
}

func main() {
	fmt.Println(
		runningSum([]int{1, 2, 3, 4}),
		runningSum([]int{1, 1, 1, 1, 1}),
		runningSum([]int{3, 1, 2, 10, 1}),
	)

}
