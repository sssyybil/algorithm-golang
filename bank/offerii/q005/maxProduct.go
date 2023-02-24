package main

import (
	"math"
	"sort"
)

// 【剑指 Offer II 005. 单词长度的最大乘积】
//	🫥https://leetcode.cn/problems/aseY1I/?envType=study-plan&id=lcof-ii&plan=lcof&plan_progress=cmwvxfi

/*
*
【一】暴力破解，找出所有的字符串组合，并判断每个组合中的两个字符串是否包含相同字符。

	❌ 超时
*/
func maxProductA(words []string) int {
	n, ans := len(words), 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if !containsDuplicateStr(words[i], words[j]) {
				ans = int(math.Max(float64(len(words[i])*len(words[j])), float64(ans)))
			}
		}
	}
	return ans
}

func containsDuplicateStr(s1, s2 string) bool {
	record := make(map[rune]bool)
	for _, v := range []rune(s1) {
		record[v] = true
	}
	for _, v := range []rune(s2) {
		if record[v] {
			return true
		}
	}
	return false
}

/*
*
【二】暴力破解优化，对原数组排序，找出所有的字符串组合，首先判断待组合的两个字符串的首字母是否相同，若相同则跳过，否则再判断每个组合中的两个字符串是否包含相同字符。

	❌ 超时
*/
func maxProductB(words []string) int {
	n, ans := len(words), 0
	sort.Strings(words)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if words[i][0] == words[j][0] {
				continue
			}
			if !containsDuplicateStr(words[i], words[j]) {
				ans = int(math.Max(float64(len(words[i])*len(words[j])), float64(ans)))
			}
		}
	}
	return ans
}

/*
*
【三】位运算
*/
func maxProductC(words []string) int {
	masks := make([]int, len(words))
	ans := 0
	for i, word := range words {
		for _, ch := range word {
			masks[i] |= 1 << (ch - 'a')
		}
	}

	for i, x := range masks {
		for j, y := range masks[:i] {
			if x&y == 0 && len(words[i])*len(words[j]) > ans {
				ans = len(words[i]) * len(words[j])
			}
		}
	}

	return ans
}
