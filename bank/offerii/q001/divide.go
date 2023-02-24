package main

import "math"

// 【剑指 Offer II 001. 整数除法】
// 🫥https://leetcode.cn/problems/xoh6Oh/description/?envType=study-plan&id=lcof-ii&plan=lcof&plan_progress=cmwvxfi

func divide(a int, b int) int {
	// 条件为「只能存储 23 位有符号整数」
	// 考虑被除数为最小值的情况
	if a == math.MinInt32 {
		if b == 1 {
			return math.MinInt32
		}
		if b == -1 {
			return math.MaxInt32
		}
	}
	// 考虑除数为最小值的情况
	if b == math.MinInt32 {
		if a == math.MinInt32 {
			return 1
		}
		return 0
	}
	// 考虑被除数为 0 的情况
	if a == 0 {
		return 0
	}

	rev := false
	if a > 0 {
		a = -a
		rev = !rev
	}
	if b > 0 {
		b = -b
		rev = !rev
	}

	ans := 0
	left, right := 1, math.MaxInt32
	for left <= right {
		mid := left + (right-left)>>1
		if quickAdd(b, mid, a) {
			ans = mid
			if mid == math.MaxInt32 {
				break
			}
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if rev {
		return -ans
	}
	return ans
}

func quickAdd(y, z, x int) bool {
	// z = z/2
	for res, add := 0, y; z > 0; z >>= 1 {
		if z&1 > 0 {
			if res < x-add {
				return false
			}
			res += add
		}
		if z != 1 {
			if add < x-add {
				return false
			}
			add += add
		}
	}
	return true
}
