package main

import (
	"strconv"
)

// 【剑指 Offer 15. 二进制中1的个数】🌀https://leetcode.cn/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/description/

func hammingWeight(num uint32) int {
	// rewrite num to binary num
	//toBinary := func(n int) string {
	//	res := ""
	//	for n != 0 {
	//		div, remain := n/2, n%2
	//		res += strconv.Itoa(remain)
	//		n = div
	//	}
	//	return res
	//}

	binaryNum := strconv.FormatInt(int64(num), 2)
	count := 0
	for _, v := range binaryNum {
		if string(v) == "1" {
			count++
		}
	}
	return count
}
