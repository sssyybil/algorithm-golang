package main

// 【面试题 01.07. 旋转矩阵】🌀https://leetcode.cn/problems/rotate-matrix-lcci/description/?favorite=xb9lfcwi

// rotate 使用「辅助数组」
func rotate(matrix [][]int) {
	n := len(matrix)
	duplicate := make([][]int, len(matrix))
	for i := range duplicate {
		duplicate[i] = make([]int, n)
	}

	for i, row := range matrix {
		for j, column := range row {
			duplicate[j][n-i-1] = column
		}
	}
	copy(matrix, duplicate)
}
