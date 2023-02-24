package main

// 【剑指 Offer II 006. 排序数组中两个数字之和】
//	🫥https://leetcode.cn/problems/kLl5u1/description/?envType=study-plan&id=lcof-ii&plan=lcof&plan_progress=cmwvxfi

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l, r}
		} else if sum < target {
			l++
		} else {
			r--
		}
	}
	return []int{}
}
