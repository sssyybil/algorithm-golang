package main

import "fmt"

/**
 * 【509. 斐波那契数】🐱https://leetcode.cn/problems/fibonacci-number/
 */

func fibWithRecur(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibWithRecur(n-1) + fibWithRecur(n-2)
}

var memory = make([]int, 31)

func fib(n int) int {
	memory[0] = 0
	memory[1] = 1
	if n == 0 || n == 1 {
		return memory[n]
	}
	res := 0
	for i := 2; i <= n; i++ {
		res = memory[i-1] + memory[i-2]
		memory[i] = res
	}
	return memory[n]
}

func main() {

	fmt.Println(
		fibWithRecur(30),
		fib(30),
	)
}
