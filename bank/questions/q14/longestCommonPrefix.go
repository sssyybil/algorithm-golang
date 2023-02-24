package main

import "fmt"

/**
 * 【14. 最长公共前缀】🐱https://leetcode.cn/problems/longest-common-prefix/
 */

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	// 变量 i 遍历列
	for i := 0; i < len(strs[0]); i++ {
		// 变量 j 遍历行
		for j := 1; j < len(strs); j++ {
			// 如果当前遍历到的列（i 值）与当前遍历的字符长度相同 -> 当前字符已经为字符数组中的最短字符，可做为结果返回 或
			// 当前遍历的字符与第一个字符串中相同位置的字符不同 -> 以字符数组中的第一个字符为基准，当前遍历字符与基准字符相同位置的字符不同，即不符合公共前缀的要求，可作为结果返回
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

func main() {
	fmt.Println(
		longestCommonPrefix([]string{"flower", "flow", "flowaht"}),
		longestCommonPrefix([]string{"flower", "flow", "flight"}),
	)
}
