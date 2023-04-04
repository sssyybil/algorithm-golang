package main

import "math"

/**
 * 【866. 回文素数】
 * 🪢https://leetcode.cn/problems/prime-palindrome/
 * @2023-03-31 14:45
 */

func primePalindrome(n int) int {
	if n == 1 || n == 2 {
		return 2
	}
	for {
		// n 若为偶数，则一定不为质数（判断 n 是否为偶数比判断 n 是否为质数更快捷），如果当前数字长度为 8，则可以跳过检查，因为不存在长度为 8 的素数
		if n%2 == 0 || (10000000 <= n && n <= 100000000) {
			n++
			continue
		}
		if isPalindrome(n) && isPrimeNumber(n) {
			return n
		}
		n++
	}
}

// isPrimeNumber 判断 n 是否为素数（素数：质数又称素数。一个大于1的自然数，除了1和它自身外，不能被其他自然数整除的数叫做质数;否则称为合数。）
func isPrimeNumber(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// isPalindrome 判断是否为回文
func isPalindrome(n int) bool {
	if n < 10 || (n%10 == 0 && n != 0) {
		return true
	}

	revertedNumber := 0
	for n > revertedNumber {
		revertedNumber = revertedNumber*10 + n%10
		n /= 10
	}
	return revertedNumber == n || revertedNumber/10 == n
}
