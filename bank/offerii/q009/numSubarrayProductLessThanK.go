package main

// 【剑指 Offer II 009. 乘积小于 K 的子数组】
//	🫥https://leetcode.cn/problems/ZVAVXX/description/

// numSubarrayProductLessThanK 暴力破解，双层嵌套循环，时间复杂度为 O(n^2)
func numSubarrayProductLessThanK(nums []int, k int) int {
	n, count := len(nums), 0
	for i := 0; i < n; i++ {
		multi := 1
		for j := i; j < n; j++ {
			multi *= nums[j]
			if multi < k {
				count++
			} else {
				break
			}
		}
	}
	return count
}

// TODO 🤔尝试其他解法，比如【二分搜索】和【滑动窗口】
