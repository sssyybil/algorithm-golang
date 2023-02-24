package main

import "fmt"

/**
 * 【1893. 检查是否区域内所有整数都被覆盖】🐱https://leetcode.cn/problems/check-if-all-the-integers-in-a-range-are-covered/
 */

func isCovered(ranges [][]int, left int, right int) bool {
	for i := left; i <= right; i++ {
		if exist(ranges, i) {
			continue
		} else {
			return false
		}
	}
	return true
}
func exist(ranges [][]int, n int) bool {
	for i := 0; i < len(ranges); i++ {
		l := ranges[i][0]
		r := ranges[i][1]
		if n >= l && n <= r {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(
		isCovered([][]int{{1, 2}, {3, 4}, {5, 6}}, 2, 5),
		isCovered([][]int{{1, 10}, {10, 20}}, 21, 21),
	)
}
