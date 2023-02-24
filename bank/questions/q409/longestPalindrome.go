package main

import "fmt"

/**
 * 【409. 最长回文串】🐱https://leetcode.cn/problems/longest-palindrome/
 */

func longestPalindrome(s string) int {
	countChar := make(map[rune]int)
	for _, ch := range []rune(s) {
		count, ok := countChar[ch]
		if ok {
			count++
			countChar[ch] = count
		} else {
			countChar[ch] = 1
		}
	}
	// ans 存储回文串的长度
	ans := 0
	for _, v := range countChar {
		ans += v / 2 * 2
		// 取一个质数个的字符作为回文串的中心
		if v%2 == 1 && ans%2 == 0 {
			ans++
		}
	}

	return ans
}

func main() {
	fmt.Println(
		longestPalindrome("aaaaakkk"),
		longestPalindrome("abccccdd"),
	)
}
