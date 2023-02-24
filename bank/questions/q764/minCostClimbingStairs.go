package main

import (
	"fmt"
	"math"
)

/**
 * 【746. 使用最小花费爬楼梯】🐱https://leetcode.cn/problems/min-cost-climbing-stairs/
 */

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = int(math.Min(float64(dp[i-1]+cost[i-1]), float64(dp[i-2]+cost[i-2])))
	}
	return dp[n]
}

func main() {
	fmt.Println(
		minCostClimbingStairs([]int{10, 15, 20}),
		minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}),
	)
}
