package main

import "fmt"

/**
 * 【54. 螺旋矩阵】🐱https://leetcode.cn/problems/spiral-matrix/
 */

func spiralOrder(matrix [][]int) []int {
	var ans []int
	upRow, downRow, leftColumn, rightColumn := 0, len(matrix)-1, 0, len(matrix[0])-1
	for {
		// 从 左->右 遍历行
		for i := leftColumn; i <= rightColumn; i++ {
			ans = append(ans, matrix[upRow][i])
		}
		upRow += 1
		if upRow > downRow {
			break
		}

		// 从 上->下 遍历列
		for i := upRow; i <= downRow; i++ {
			ans = append(ans, matrix[i][rightColumn])
		}
		rightColumn -= 1
		if rightColumn < leftColumn {
			break
		}

		// 从 右->左 遍历行
		for i := rightColumn; i >= leftColumn; i-- {
			ans = append(ans, matrix[downRow][i])
		}
		downRow -= 1
		if downRow < upRow {
			break
		}

		// 从 下->上 遍历列
		for i := downRow; i >= upRow; i-- {
			ans = append(ans, matrix[i][leftColumn])
		}
		leftColumn += 1
		if leftColumn > rightColumn {
			break
		}
	}
	return ans
}

func main() {
	fmt.Println(
		spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}),
		spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}),
	)
}
