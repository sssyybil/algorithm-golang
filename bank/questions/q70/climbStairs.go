package main

import "fmt"

/**
 * 【70. 爬楼梯】🐱https://leetcode.cn/problems/climbing-stairs/
 */

func climbStairs(n int) int {
	memory := []int{1, 1}
	if n <= 1 {
		return memory[n]
	}
	sum := 0
	for i := 2; i <= n; i++ {
		sum = memory[i-1] + memory[i-2]
		memory = append(memory, sum)
	}
	return memory[n]
}

func main() {

	fmt.Println(
		climbStairs(10),
	)

}
