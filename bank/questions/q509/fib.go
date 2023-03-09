package main

/**
 * 【509. 斐波那契数】🐱https://leetcode.cn/problems/fibonacci-number/
 */

// fibWithRecur 递归，使用【自上而下】的解决方式，即假设最小的问题已经解决了。时间复杂度呈指数级增长。
func fibWithRecur(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibWithRecur(n-1) + fibWithRecur(n-2)
}

var record []int

// fibWithRecurBetter 使用【记忆搜索法】优化递归求解的过程，时间复杂度为 O(n)
func fibWithRecurBetter(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if record[n] == 0 {
		record[n] = fibWithRecurBetter(n-1) + fibWithRecurBetter(n-2)
	}
	return record[n]
}

// fib 使用【自下而上】的解决方式，即先从最小的单位开始求解。通常这个过程被称为【动态规划】。时间复杂度为 O(n)
func fib(n int) int {
	// 此处声明 memo 数组的容量为 31，是因为题目中给出了 n 的范围是 0 <= n <= 30
	var memo = make([]int, 31)
	memo[0] = 0
	memo[1] = 1

	if n == 0 || n == 1 {
		return memo[n]
	}
	for i := 2; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}
	return memo[n]
}
