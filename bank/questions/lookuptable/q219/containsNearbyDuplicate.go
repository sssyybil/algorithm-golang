package main

import "fmt"

/**
*
* 【219. 存在重复元素 II】
🌀https://leetcode.cn/problems/contains-duplicate-ii/description/
*
* @author  sun
* @date 2022/11/6 11:02
*/
// 使用查找表
func containsNearbyDuplicateA(nums []int, k int) bool {
	table := make(map[int]int)
	for i, v := range nums {
		if curIndex, has := table[v]; has && i-curIndex <= k {
			return true
		}
		table[v] = i
	}

	return false
}

// 滑动窗口
func containsNearbyDuplicateB(nums []int, k int) bool {
	record := make(map[int]bool)
	for i, v := range nums {
		// 保证 record 中最多有 k 个元素
		if i > k {
			delete(record, nums[i-k-1])
		}
		if record[v] {
			return true
		}
		record[v] = true
	}
	return false
}

func main() {
	fmt.Println(
		containsNearbyDuplicateA([]int{99, 99}, 2),           // true
		containsNearbyDuplicateA([]int{1, 2, 3, 1}, 3),       // true
		containsNearbyDuplicateA([]int{1, 0, 1, 1}, 1),       // true
		containsNearbyDuplicateA([]int{1, 2, 3, 1, 2, 3}, 2), // false
	)

	fmt.Println(
		containsNearbyDuplicateB([]int{99, 99}, 2),           // true
		containsNearbyDuplicateB([]int{1, 2, 3, 1}, 3),       // true
		containsNearbyDuplicateB([]int{1, 0, 1, 1}, 1),       // true
		containsNearbyDuplicateB([]int{1, 2, 3, 1, 2, 3}, 2), // false
	)
}
