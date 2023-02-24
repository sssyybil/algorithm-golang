package main

import (
	"fmt"
	"unicode"
)

/**
 * 【1417. 重新格式化字符串】🐱https://leetcode.cn/problems/reformat-the-string/submissions/
 */

func reformat(s string) string {
	var sliceDigit []string
	var sliceCharacter []string
	for _, ch := range s {
		if unicode.IsDigit(ch) {
			sliceDigit = append(sliceDigit, string(ch))
		} else {
			sliceCharacter = append(sliceCharacter, string(ch))
		}
	}
	n, m := len(sliceDigit), len(sliceCharacter)
	if abs(n-m) > 1 {
		return ""
	}
	res, total := "", m+n
	for len(res) != total {
		if n > m {
			n--
			res += sliceDigit[n]
		} else if n < m {
			m--
			res += sliceCharacter[m]
		} else {
			if len(res) > 0 && unicode.IsDigit(rune(res[len(res)-1])) {
				m--
				res += sliceCharacter[m]
			} else {
				n--
				res += sliceDigit[n]
			}
		}
	}
	return res
}

func abs(v int) int {
	if v < 0 {
		return -v
	} else {
		return v
	}
}

func main() {
	fmt.Println(
		reformat("covid2019"),
		reformat("aaa456"),
		reformat("leetcode"),
	)
}
