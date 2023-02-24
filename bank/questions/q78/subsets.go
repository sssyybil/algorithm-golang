package main

import (
	"fmt"
)

/**
 * 【1417. 重新格式化字符串】🐱https://leetcode.cn/problems/reformat-the-string/submissions/
 */

// TODO repeat

func subsets(nums []int) [][]int {
	var ans [][]int
	n := len(nums)
	for mask := 0; mask < 1<<n; mask++ {
		var set []int
		for i, v := range nums {
			if mask>>i&1 > 0 {
				set = append(set, v)
			}
		}
		ans = append(ans, append([]int(nil), set...))
	}
	return ans
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}
