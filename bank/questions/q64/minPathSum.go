package main

import "fmt"

/**
 * 【64. 最小路径和】🐱https://leetcode.cn/problems/minimum-path-sum/description/
 */

func minPathSum(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	matrix := make([][]int, row)
	for i := 0; i < row; i++ {
		matrix[i] = make([]int, col)
	}
	matrix[0][0] = grid[0][0]

	// 初始化行的值
	for r := 1; r < row; r++ {
		matrix[r][0] = matrix[r-1][0] + grid[r][0]
	}
	// 初始化列的值
	for c := 1; c < col; c++ {
		matrix[0][c] = matrix[0][c-1] + grid[0][c]
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			matrix[i][j] = min(matrix[i-1][j], matrix[i][j-1]) + grid[i][j]
		}
	}

	return matrix[row-1][col-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
	fmt.Println(minPathSum([][]int{{1, 2, 3}, {4, 5, 6}}))
}
