package main

import "strconv"

/**
 * 【9. 回文数】
 * 🪢https://leetcode.cn/problems/palindrome-number/description/
 * @2023-03-31 12:19
 */

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x >= 0 && x < 10 {
		return true
	}
	str := strconv.Itoa(x)
	l, r := 0, len(str)-1
	for l < r {
		if str[l] != str[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func isPalindromeWithoutStr(x int) bool {
	// 当 x 为负数时不可能为回文，或者当 x 最后一位为 0 时，若想满足 x 为回文，则 x==0，反之，若 x 最后一位为 0 且 x != 0，则 x 不是回文
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}
	return revertedNumber == x || revertedNumber/10 == x
}
