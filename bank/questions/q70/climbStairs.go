package main

/**
 * 【70. 爬楼梯】🐱https://leetcode.cn/problems/climbing-stairs/
 */

// climbStairs 使用【动态规划】求解
func climbStairs(n int) int {
	memory := []int{1, 1}
	if n <= 1 {
		return memory[n]
	}
	sum := 0
	for i := 2; i <= n; i++ {
		sum = memory[i-1] + memory[i-2]
		memory = append(memory, sum)
	}
	return memory[n]
}

// climbStairsWithMap 动态规划求解，与 climbStairs 思路相同
func climbStairsWithMap(n int) int {
	memo := make(map[int]int)
	memo[1] = 1
	memo[2] = 2
	for i := 3; i <= n; i++ {
		if v, has := memo[i]; has {
			return v
		}
		memo[i] = memo[i-1] + memo[i-2]
	}
	return memo[n]
}

// climbStairsWithRecur 使用递归方式求解
func climbStairsWithRecur(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairsWithRecur(n-1) + climbStairsWithRecur(n-2)
}

var memo = make(map[int]int)

// climbStairsWithMemo 使用【记忆化搜索法】优化递归求解
func climbStairsWithMemo(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if memo[n] == 0 {
		memo[n] = climbStairsWithMemo(n-1) + climbStairsWithMemo(n-2)
	}
	return memo[n]
}
