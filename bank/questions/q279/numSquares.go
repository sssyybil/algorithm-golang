package main

import (
	"fmt"
	"math"
)

var memo = make(map[int]int)

func numSquaresWithRecur(n int) int {
	if n == 1 {
		return 1
	}
	if memo[n] == 0 {
		res := -1
		for i := 1; i*i <= n; i++ {
			res = min(res, memo[n-i*i])
		}
		memo[n] = res + 1
	}
	return memo[n]
}

// f[0] = 0
// f[1] = 1
// f[2] = 2
// f[3] = 3
// f[4] = 1
// f[5] = 2
// f[6] = 3
// f[7] = 4
// f[8] = 2
func numSquares(n int) int {
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minN := math.MaxInt32
		// [1, 2, 4, 9, 16]
		for j := 1; j*j <= i; j++ {
			minN = min(minN, f[i-j*j])
		}
		f[i] = minN + 1
	}
	for i, v := range f {
		fmt.Printf("f[%v] = %v\n", i, v)
	}
	return f[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
