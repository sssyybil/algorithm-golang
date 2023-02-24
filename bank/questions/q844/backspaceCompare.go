package main

import (
	"fmt"
)

/**
 * 【844. 比较含退格的字符串】🐱https://leetcode.cn/problems/backspace-string-compare/
 */

func backspaceCompare(s string, t string) bool {
	return rebuildStr(s) == rebuildStr(t)
}

func rebuildStr(s string) string {
	var res []rune
	for _, v := range []rune(s) {
		if string(v) != "#" {
			res = append(res, v)
		} else if len(res) > 0 {
			res = res[:len(res)-1]
		}
	}
	return string(res)
}

func main() {

	fmt.Println(
		backspaceCompare("a##c", "#a#c"),
		backspaceCompare("xywrrmp", "xywrrmu#p"),
		backspaceCompare("abccc#i", "abccd#i"),
		backspaceCompare("ab#c", "ad#c"),
		backspaceCompare("ab##", "c#d#"),
		backspaceCompare("ac#", "c"),
	)

	s := "abc"
	arr := make([]rune, 6)
	for i, v := range []rune(s) {
		arr[i] = v
	}
	fmt.Printf("%c", arr)
	fmt.Println(arr)

}
