package main

/**
 * 【1480. 一维数组的动态和】🐱https://leetcode.cn/problems/running-sum-of-1d-array/
 */

func runningSum(nums []int) []int {
	// res 声明成数组类型要比声明成切片类型快很多
	res, curSum := make([]int, len(nums)), 0
	for i, v := range nums {
		curSum += v
		res[i] = curSum
	}
	return res
}
