package main

import "fmt"

/**
*
* 【1. 两数之和】
🐱https://leetcode.cn/problems/two-sum/
*
* @author  sun
* @date 2022/11/1 09:58
*/

// 使用「查找表」求解，时间复杂度为 O(N)；空间复杂度为 O(N)，其中 N 为数组中元素的个数
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		diff := target - v
		if index, has := m[diff]; has {
			return []int{index, i}
		}
		m[v] = i
	}
	return []int{}
}

func main() {
	fmt.Println(
		twoSum([]int{2, 7, 11, 15}, 9), // [0,1]
		twoSum([]int{3, 2, 4}, 6),      // [1,2]
		twoSum([]int{3, 3}, 6),         // [0,1]
	)
}
