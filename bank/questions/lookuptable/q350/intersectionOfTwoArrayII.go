package main

import "fmt"

/**
 * 【350. 两个数组的交集 II 】🐱https://leetcode.cn/problems/intersection-of-two-arrays-ii/
 */

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _, n1 := range nums1 {
		if v, has := m[n1]; has {
			m[n1] = v + 1
		} else {
			m[n1] = 1
		}
	}

	var ans []int
	for _, n2 := range nums2 {
		v, has := m[n2]
		if has && v > 0 {
			ans = append(ans, n2)
			m[n2] = v - 1
		}
	}
	return ans
}

func main() {
	fmt.Println(
		intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}),
		intersect([]int{1, 2, 2, 1}, []int{2, 2}),
		intersect([]int{1, 2, 3, 4, 5, 6, 6}, []int{2, 3}),
	)
}
