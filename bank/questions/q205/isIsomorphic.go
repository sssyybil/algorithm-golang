package main

import "fmt"

/**
 * 【205. 同构字符串】🐱https://leetcode.cn/problems/isomorphic-strings/
 */

// 本题需要判断 s 和 t 中的每个字符都一一对应，即 s 中的任意字符被 t 中的唯一字符对应，t 中的任意字符被 s 中的字符唯一对应，这种关系也被称为『双射』。
// 题目条件中已经给出：s.length == t.length
func isIsomorphic(s string, t string) bool {
	s2t := make(map[byte]byte)
	t2s := make(map[byte]byte)
	for i := range s {
		a, b := s[i], t[i]
		if (s2t[a] > 0 && s2t[a] != b) || (t2s[b] > 0 && t2s[b] != a) {
			return false
		}
		s2t[a] = b
		t2s[b] = a
	}
	return true
}

func main() {
	fmt.Println(
		isIsomorphic("badc", "baba"),
		isIsomorphic("egg", "add"),
		isIsomorphic("abd", "dea"),
		isIsomorphic("foo", "bar"),
		isIsomorphic("paper", "title"),
		isIsomorphic("a", "c"),
	)
}
