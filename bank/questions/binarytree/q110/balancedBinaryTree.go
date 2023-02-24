package main

import (
	"algorithm-golang/types"
	"fmt"
	"math"
)

/**
 * 【110. 平衡二叉树】🐱https://leetcode.cn/problems/balanced-binary-tree/
 */

func isBalanced(root *types.TreeNode) bool {
	return recur(root) != -1
}

func recur(root *types.TreeNode) int {
	if root == nil {
		return 0
	}
	left := recur(root.Left)
	if left == -1 {
		return -1
	}
	right := recur(root.Right)
	if right == -1 {
		return -1
	}

	// 若左子树和右子树的高度差 < 2，则返回以 root 为根节点的子树的最大高度
	if int(math.Abs(float64(left-right))) < 2 {
		return int(math.Max(float64(left), float64(right))) + 1
	}
	return -1
}

func main() {
	root := types.TreeNode{Val: 1}
	root.Left = &types.TreeNode{Val: 2}
	root.Right = &types.TreeNode{Val: 3}
	root.Left.Left = &types.TreeNode{Val: 4}
	root.Left.Right = &types.TreeNode{Val: 5}
	root.Right.Left = &types.TreeNode{Val: 6}
	fmt.Println(isBalanced(&root))
}
