package main

import (
	"fmt"
)

/**
 * 【22. 括号生成】🐱https://leetcode.cn/problems/generate-parentheses/
 */

// TODO repeat

func generateParenthesis(n int) []string {
	var combinations []string
	generateAll(make([]rune, 2*n), 0, &combinations)
	return combinations
}

func generateAll(current []rune, pos int, result *[]string) {
	if pos == len(current) {
		if valid(current) {
			*result = append(*result, string(current))
			fmt.Printf("Append Happend.current result = %v\n", result)
		}
	} else {
		current[pos] = '('
		generateAll(current, pos+1, result)
		current[pos] = ')'
		generateAll(current, pos+1, result)
	}
}

func valid(current []rune) bool {
	balance := 0
	for _, v := range current {
		if v == '(' {
			balance++
		} else {
			balance--
		}
		// ⚠️ 判断括号是否有效，这行非常重要
		if balance < 0 {
			return false
		}
	}
	return balance == 0
}

func main() {
	str := "(()))("
	fmt.Println(
		valid([]rune(str)),
		generateParenthesis(3),
	)

}
