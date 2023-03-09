package main

import (
	"math"
)

/**
 * 【120. 三角形最小路径和】🐱https://leetcode.cn/problems/triangle/description/
 *
	==> 题目中给出：-104 <= triangle[i][j] <= 104，因此需要考虑元素存在负数的情况
*/

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	memo := make([][]int, n)
	for i, v := range triangle {
		memo[i] = make([]int, len(v))
	}
	memo[0][0] = triangle[0][0]

	// ⚠️此处需要注意的是使用数组构成的三角形的特性，第一行一个元素，第二行两个元素，第三行三个元素……以此类推
	for i := 1; i < len(triangle); i++ {
		memo[i][0] = memo[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			memo[i][j] = min(memo[i-1][j-1], memo[i-1][j]) + triangle[i][j]
		}
		memo[i][i] = memo[i-1][i-1] + triangle[i][i]
	}

	res := math.MaxInt32
	for i := 0; i < n; i++ {
		res = min(memo[n-1][i], res)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
