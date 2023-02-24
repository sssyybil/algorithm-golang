package main

// 【剑指 Offer II 004. 只出现一次的数字 】
// 🫥https://leetcode.cn/problems/WGki4K/description/?envType=study-plan&id=lcof-ii&plan=lcof&plan_progress=cmwvxfi

func singleNumber(nums []int) int {
	record := make(map[int]int)
	for _, v := range nums {
		record[v]++
	}
	for k, v := range record {
		if v == 1 {
			return k
		}
	}
	return -1
}
