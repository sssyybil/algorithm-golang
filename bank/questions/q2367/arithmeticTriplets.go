package main

/**
 * 【2367. 算术三元组的数目】
 * 🪢https://leetcode.cn/problems/number-of-arithmetic-triplets/description/
 * @2023-03-31 11:45
 */
func arithmeticTriplets(nums []int, diff int) int {
	n, res := len(nums), 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[j]-nums[i] == diff {
				for k := j + 1; k < n; k++ {
					if nums[k]-nums[j] == diff {
						res++
					}
				}
			}
		}
	}
	return res
}
