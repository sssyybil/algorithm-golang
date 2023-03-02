package main

/**
 * 【3. 无重复字符的最长子串】🐱https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/
 */

// lengthOfLongestSubstring 暴力破解，超时
func lengthOfLongestSubstring(s string) int {
	n, res := len(s), 0
	if n == 1 {
		return n
	}
	//record := make(map[string]int)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			cur := s[i : j+1]
			if !isDuplicate(cur) && len(cur) > res {
				res = len(cur)
			}
		}
	}

	return res
}

func isDuplicate(s string) bool {
	record := make(map[rune]int)
	for _, r := range []rune(s) {
		if _, has := record[r]; has {
			return true
		}
		record[r] = 1
	}
	return false
}

// lengthOfLongestSubstringB 滑动窗口
func lengthOfLongestSubstringB(s string) int {
	n, res := len(s), 0
	// 当 s ==" " 时，符合题，结果应返回 1
	if n == 1 {
		return n
	}
	l, r := 0, 0
	// 滑动窗口中的 r 指针结束滑动的原因：在当前的窗口中遇到了已经存在的元素 或 当前元素已经遍历完成且全部为不重复的元素
	for l < n {
		recode := make(map[string]bool)
		for r = l; r < n; r++ {
			if recode[string(s[r])] {
				break
			}
			recode[string(s[r])] = true
		}
		// 计算当前不重复子串的最大长度时，应根据 record 的长度来计算，而不是在指针 r 遍历过程中，r 指针遇到重复元素时候与 l 指针的距离
		if len(recode) > res {
			res = len(recode)
		}
		l++
	}

	return res
}
