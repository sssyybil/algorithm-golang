package main

import "fmt"

/**
*
* 【1684. 统计一致字符串的数目】
🌀https://leetcode.cn/problems/count-the-number-of-consistent-strings/description/
*
* @author  sun
* @date 2022/11/8 10:06
*/
// 使用额外的存储空间 map ➕两次嵌套循环
func countConsistentStringsA(allowed string, words []string) int {
	record := make(map[rune]bool)
	for _, v := range []rune(allowed) {
		record[v] = true
	}

	count := 0
	for _, word := range words {
		contains := false
		for _, s := range []rune(word) {
			if !record[s] {
				contains = false
				break
			}
			contains = true
		}
		if contains {
			count++
		}
	}
	return count
}

func countConsistentStringsB(allowed string, words []string) int {
	mask := 0
	for _, c := range allowed {
		mask |= 1 << (c - 'a')
	}

	return 0
}

func main() {
	fmt.Println(
		countConsistentStringsA("ab", []string{"ad", "bd", "aaab", "baa", "badab"}),               // 2，["aaab", "baa"]
		countConsistentStringsA("abc", []string{"a", "b", "c", "ab", "ac", "bc", "abc"}),          // 7，["a", "b", "c", "ab", "ac", "bc", "abc"]
		countConsistentStringsA("cad", []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}), // 4，["cc", "acd","ac", "d"]
	)

	fmt.Printf("%b\n", 12)

}
