package main

import "fmt"

/**
 * 【392. 判断子序列】🐱https://leetcode.cn/problems/is-subsequence/
 */

func isSubsequence(s string, t string) bool {
	n1, n2 := len(s), len(t)
	if n1 == 0 {
		return true
	}
	if n1 > n2 {
		return false
	}
	i, j := 0, 0
	for j = range t {
		if i == n1 {
			break
		}
		if s[i] == t[j] {
			i++
		}
	}
	return i == n1
}

func main() {

	fmt.Println(
		isSubsequence("abc", "ahbgdc"),
		isSubsequence("axc", "ahbgdc"),
		isSubsequence("aabb", "aaabbb"),
	)

}
