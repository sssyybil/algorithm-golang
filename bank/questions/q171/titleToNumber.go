package main

import (
	"math"
)

/**
 * 【171. Excel 表列序号】🐱https://leetcode.cn/problems/excel-sheet-column-number/description/
 */

// 相当于使用 26 机制去计算，ABD = 1*26^2 + 2*26^1 + 4*26^0 = 732
func titleToNumber(columnTitle string) int {
	n := len(columnTitle)
	if n == 1 {
		return int(columnTitle[0] - 'A' + 1)
	}
	ans := 0
	for _, v := range []rune(columnTitle) {
		num := int(v - 'A' + 1)
		n = n - 1
		ans += num * int(math.Pow(26, float64(n)))
	}
	return ans
}
