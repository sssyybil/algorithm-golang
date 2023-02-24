package main

import (
	"fmt"
	"strconv"
)

/**
 * 【202. 快乐数】🐱https://leetcode.cn/problems/happy-number/
 */

func isHappy(n int) bool {
	memory := make(map[int]int)
	res := n
	for {
		if res == 1 {
			return true
		} else {
			res = getNext(res)
			if _, ok := memory[res]; !ok {
				memory[res] = 1
			} else {
				return false
			}
		}
	}
}

func getNext(num int) int {
	str, sum := strconv.Itoa(num), 0
	for i := 0; i < len(str); i++ {
		remainder := num % 10
		sum += remainder * remainder
		num = num / 10
	}
	return sum
}

func main() {
	fmt.Println(
		isHappy(19),
		isHappy(123),
		isHappy(2),
	)
}
