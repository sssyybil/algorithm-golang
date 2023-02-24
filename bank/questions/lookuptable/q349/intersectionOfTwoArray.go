package main

import "fmt"

/**
 * 【349. 两个数组的交集】🐱https://leetcode.cn/problems/intersection-of-two-arrays/
 */

// intersection 在 Go 中，没有 set 这种数据结构，但是可以使用 map 实现 set 相应的功能
func intersection(nums1 []int, nums2 []int) []int {
	var ans []int
	m := make(map[int]bool)
	for _, v := range nums1 {
		m[v] = true
	}

	for _, v := range nums2 {
		if m[v] {
			ans = append(ans, v)
			m[v] = false
		}
	}
	return ans
}

func main() {
	fmt.Println(
		intersection([]int{1, 2, 2, 1}, []int{2, 2}),
		intersection([]int{1, 2, 3, 4, 5, 6, 6}, []int{2, 3}),
	)
}
