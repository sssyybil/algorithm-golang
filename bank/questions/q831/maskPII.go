package main

import (
	"strings"
)

/**
 * 【831. 隐藏个人信息】
 * 🪢https://leetcode.cn/problems/masking-personal-information/
 * @2023-04-01 10:28
 */

func maskPII(s string) string {
	if strings.ContainsAny(s, "@") {
		return encryptEmail(s)
	}

	return encryptPhone(s)
}

func encryptEmail(s string) string {
	lower := strings.ToLower(s)

	split := strings.Split(lower, "@")
	emailName := split[0]
	return string(emailName[0]) + "*****" + string(emailName[len(emailName)-1]) + "@" + split[1]
}

func encryptPhone(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "+", "")
	s = strings.ReplaceAll(s, "-", "")
	s = strings.ReplaceAll(s, "(", "")
	s = strings.ReplaceAll(s, ")", "")
	s = strings.ReplaceAll(s, " ", "")

	n, pre := len(s), s[:len(s)-4]
	localNumber := "***-***-" + s[len(pre):]
	if n == 10 {
		return localNumber
	}

	res := "+"
	for i := 0; i < n-10; i++ {
		res += "*"
	}

	return res + "-" + localNumber
}
