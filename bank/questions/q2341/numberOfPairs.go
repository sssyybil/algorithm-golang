package main

/**
 * 【2341. 数组能形成多少数对】🐱https://leetcode.cn/problems/maximum-number-of-pairs-in-array/
 */

func numberOfPairs(nums []int) []int {
	record := make(map[int]int)
	pair, residue := 0, 0
	for _, v := range nums {
		if c, has := record[v]; has {
			record[v] = c + 1
		} else {
			record[v] = 1
		}
	}

	for _, v := range record {
		if v == 2 {
			pair++
		} else {
			ans := v / 2
			pair += ans
			if v%2 != 0 {
				residue++
			}
		}
	}
	return []int{pair, residue}
}
