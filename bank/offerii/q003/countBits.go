package main

import (
	"math/bits"
	"strconv"
)

// 【剑指 Offer II 003. 前 n 个数字二进制中 1 的个数】
//	🫥https://leetcode.cn/problems/w3tCBm/?envType=study-plan&id=lcof-ii&plan=lcof&plan_progress=cmwvxfi

/*
*
【一】将每个数字转成二进制字符串，然后遍历每个字符串，获得其中 1 的个数
*/
func countBitsA(n int) []int {
	ans := make([]int, n+1)
	for i := 0; i <= n; i++ {
		ans[i] = countZero(i)
	}
	return ans
}

func countZero(n int) int {
	var res string
	ans := 0
	for n > 0 {
		div, remainder := n/2, n%2
		n = div
		res += strconv.Itoa(remainder)
		if remainder == 1 {
			ans++
		}
	}
	return ans
}

/*
*
【二】使用内置函数 bits.OnesCount(x uint)
*/
func countBitsB(n int) []int {
	ans := make([]int, n+1)
	for i := 0; i <= n; i++ {
		ans[i] = bits.OnesCount(uint(i))
	}
	return ans
}

/*
*
【三】Brian Kernighan 算法：对于任意整数 x，另 x = x&(x-1)，该运算将 x 的二进制表示的最后一位 1 变成 0。
因此，对 x 重复该操作，直到 x 变成 0。
最直观的做法是对从 0 到 n 的每个整数直接计算「一比特数」。每个 int 型的数都可以用 32 位二进制数表示，只要遍历其二进制表示的每一位即可得到 1 的数目。
对于给定的 n，计算从 0 到 n 的每个整数的「一比特数」的时间都不会超过 O(logn)，因此总时间复杂度为 O(nlogn)。
*/
func countBitsC(n int) []int {
	ans := make([]int, n+1)
	for i := range ans {
		ans[i] = onesCount(i)
	}
	return ans
}

func onesCount(n int) (ones int) {
	for ; n > 0; n &= n - 1 {
		ones++
	}
	return
}
