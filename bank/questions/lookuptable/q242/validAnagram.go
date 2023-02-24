package main

import (
	"fmt"
	"sort"
	"strings"
)

/**
*
* 【242. 有效的字母异位词】
* 🔗https://leetcode.cn/problems/valid-anagram/
*
* @author  sun
* @date 2022/10/30 15:11
 */

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	ms := storeToMap(s)
	mt := storeToMap(t)
	for k, v := range ms {
		if tv, has := mt[k]; has {
			if tv != v {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

// 方式二：只统计一个字符串中各个字符的个数
func isAnagramB(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := storeToMap(s)
	for _, k := range []rune(t) {
		if sv, has := m[k]; has {
			if sv == 0 {
				return false
			} else {
				m[k] = sv - 1
			}
		} else {
			return false
		}
	}
	return true
}

// storeToMap s 仅包含小写字母
func storeToMap(s string) map[rune]int {
	m := make(map[rune]int)
	for _, c := range []rune(s) {
		if v, has := m[c]; has {
			m[c] = v + 1
		} else {
			m[c] = 1
		}
	}
	return m
}

// 方式三：排序 ➕双指针
func isAnagramC(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sc := strings.Split(s, "")
	tc := strings.Split(t, "")
	sort.Strings(sc)
	sort.Strings(tc)

	i, j := 0, 0
	for i < len(s) {
		if sc[i] != tc[j] {
			return false
		}
		i++
		j++
	}
	return true
}

func main() {
	fmt.Println(
		isAnagram("a", "ab"),            // false
		isAnagram("anagram", "nagaram"), // true
		isAnagram("rat", "car"),         // false
		isAnagram("ABC", "abc"),         // 本题目中设置 s 和 t 仅包含小写字母，相同字母的大小写格式不同，rune 值便不同。故为 false。
		isAnagram("喵喵喵", "喵喵喵汪"),        // true
	)
	fmt.Println(
		isAnagramB("a", "ab"),            // false
		isAnagramB("anagram", "nagaram"), // true
		isAnagramB("rat", "car"),         // false
		isAnagramB("ABC", "abc"),         // 本题目中设置 s 和 t 仅包含小写字母，相同字母的大小写格式不同，rune 值便不同。故为 false。
		isAnagramB("喵喵喵", "喵喵喵汪"),        // true
	)
	fmt.Println(
		isAnagramC("a", "ab"),            // false
		isAnagramC("anagram", "nagaram"), // true
		isAnagramC("rat", "car"),         // false
		isAnagramC("ABC", "abc"),         // 本题目中设置 s 和 t 仅包含小写字母，相同字母的大小写格式不同，rune 值便不同。故为 false。
		isAnagramC("喵喵喵", "喵喵喵汪"),        // true
	)
}
