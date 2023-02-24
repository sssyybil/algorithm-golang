package main

import (
	"fmt"
	"sort"
)

/**
*
* 【217. 存在重复元素】
🌀https://leetcode.cn/problems/contains-duplicate/description/
*
* @author  sun
* @date 2022/11/6 11:54
*/

func containsDuplicateA(nums []int) bool {
	record := make(map[int]bool)
	for _, v := range nums {
		if record[v] {
			return true
		}
		record[v] = true
	}
	return false
}

func containsDuplicateB(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(
		containsDuplicateA([]int{1, 2, 3, 1}),                   // true
		containsDuplicateA([]int{1, 2, 3, 4}),                   // false
		containsDuplicateA([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}), // true
	)

	fmt.Println(
		containsDuplicateB([]int{1, 2, 3, 1}),                   // true
		containsDuplicateB([]int{1, 2, 3, 4}),                   // false
		containsDuplicateB([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}), // true
	)
}
