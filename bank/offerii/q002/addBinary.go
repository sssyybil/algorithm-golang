package main

import (
	"strconv"
)

// 【剑指 Offer II 002. 二进制加法】
//	🫥https://leetcode.cn/problems/JFETK5/description/?envType=study-plan&id=lcof-ii&plan=lcof&plan_progress=cmwvxfi

func addBinary(a string, b string) string {
	n1, n2 := len(a), len(b)
	if n1 < n2 {
		return addBinary(b, a)
	}
	diff := n1 - n2
	for i := 0; i < diff; i++ {
		b = "0" + b
	}

	res, carry := "", 0
	for i := len(a) - 1; i >= 0; i-- {
		n1, err := strconv.Atoi(string(a[i]))
		if err != nil {
			panic(err)
		}
		n2, err := strconv.Atoi(string(b[i]))
		if err != nil {
			panic(err)
		}
		sum := n1 + n2 + carry
		if sum == 3 {
			res = "1" + res
			carry = 1
		} else if sum == 2 {
			res = "0" + res
			carry = 1
		} else {
			res = strconv.Itoa(sum) + res
			carry = 0
		}
	}
	if carry == 1 {
		res = "1" + res
	}

	return res
}
