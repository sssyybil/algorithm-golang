package main

import "fmt"

// 【二分查找法】，在有序数组 nums 中查找 target，如果找到 target，则返回相应的索引 index，如果没有，则返回 -1
// ==> ⚠️ 本函数适用于数组中不包含重复元素，当然，数组中包含重复元素的话，也是能找到该元素的索引位置的
// ==> 但是无法找到重复元素出现的第一个索引位置或最后一个索引位置
func search(nums []int, target int) int {
	// 在 nums[left...right] 之中查找 target
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		// 在 nums[left...mid-1] 之中查找 target
		if nums[mid] > target {
			right = mid - 1

			// 在 nums[mid+1...right] 之中查找 target
		} else {
			left = mid + 1
		}
	}
	return -1
}

// TODO 不严格递增的数组（即包含重复元素），查找 target 是否存在于该数组中，如果存在，返回 target 出现的第一个索引位置

func main() {
	fmt.Println(
		search([]int{1, 2, 4, 6, 88, 999}, 7),
		search([]int{1, 2, 4, 6, 88, 999}, 88),
		search([]int{3, 3, 3, 3, 3, 3}, 3),
	)
}
