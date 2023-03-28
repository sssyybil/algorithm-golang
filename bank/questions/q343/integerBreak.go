package main

/**
 * 【343. 整数拆分】🐱https://leetcode.cn/problems/integer-break/description/
 */

// memo 表示将数字 n 分割（至少分割成两部分）后得到的最大乘积
var memo = make(map[int]int)

// integerBreak 使用递归的方式求解，即【自顶向下】的解决方式。
// 简单思路图：https://user-images.githubusercontent.com/126223217/224233243-22d489cf-5530-4c45-a954-3d6d9b9b7231.png
func integerBreak(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	if memo[n] == 0 {
		res := -1
		// 将 n 分割成 i 和 (n-i) 两部分
		for i := 1; i < n; i++ {
			// 判断将 n 分割成 i*(n-i) 的值与再继续将 (n-i) 进行分割后的乘积值相比，哪个更大
			res = maxOfThree(res, i*(n-i), i*integerBreak(n-i))
		}
		memo[n] = res
	}
	return memo[n]
}

// integerBreakWithDp 使用动态规划的方式求解，即【自底向上】的解决方式。
func integerBreakWithDp(n int) int {
	record := make(map[int]int)
	record[1] = 1
	record[2] = 1

	for i := 3; i <= n; i++ {
		// 求解 record[i]
		for j := 1; j <= i-1; j++ {
			record[i] = maxOfThree(record[i], j*(i-j), j*record[i-j])
		}
	}
	return record[n]
}

func maxOfThree(a, b, c int) int {
	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	return max
}
