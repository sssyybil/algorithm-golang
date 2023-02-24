package main

import "fmt"

/**
*
* 【1662. 检查两个字符串数组是否相等】
🌀https://leetcode.cn/problems/check-if-two-string-arrays-are-equivalent/
*
* @author  sun
* @date 2022/11/1 10:23
*/

// 拼接字符串再进行对比
func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	// 如果 word1 的第一个字符与 word2 的第一个字符不同，则这两个字符串数组必然不同，就没有往下继续执行的必要了
	if word1[0][:1] != word2[0][:1] {
		return false
	}
	s1, s2 := "", ""
	for _, v := range word1 {
		s1 += v
	}
	for _, v := range word2 {
		s2 += v
	}
	return s1 == s2
}

func main() {
	fmt.Println(
		arrayStringsAreEqual([]string{"ab", "c"}, []string{"a", "bc"}),           // true
		arrayStringsAreEqual([]string{"a", "cb"}, []string{"ab", "c"}),           // false
		arrayStringsAreEqual([]string{"abc", "d", "defg"}, []string{"abcddefg"}), // true
	)
}
