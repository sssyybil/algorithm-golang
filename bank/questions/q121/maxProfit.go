package main

import (
	"math"
)

/**
 * 【121. 买卖股票的最佳时机】🐱https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
 */

// 双层嵌套循环，时间复杂度为 O(n^2)，超时了 ⛈
func maxProfitBad(prices []int) int {
	n, profit := len(prices), 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			profit = int(math.Max(float64(prices[j]-prices[i]), float64(profit)))
		}
	}
	return profit
}

// 一次遍历
func maxProfit(prices []int) int {
	minPrice, maxProfit := math.MaxInt, 0
	for _, v := range prices {
		minPrice = int(math.Min(float64(v), float64(minPrice)))
		maxProfit = int(math.Max(float64(v-minPrice), float64(maxProfit)))
	}
	return maxProfit
}
