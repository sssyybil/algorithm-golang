package binarytree

import (
	"algorithm-golang/types"
	"math"
)

// minDepth 求二叉树的最小深度（🎩https://leetcode.cn/problems/minimum-depth-of-binary-tree/）
func minDepth(root *types.TreeNode) int {
	if root == nil {
		return 0
	}

	minLeftDepth := minDepth(root.Left)
	minRightDepth := minDepth(root.Right)

	return int(math.Min(float64(minLeftDepth), float64(minRightDepth))) + 1
}

// maxDepth 求二叉树的最大深度（🎩https://leetcode.cn/problems/maximum-depth-of-binary-tree/）
func maxDepth(root *types.TreeNode) int {
	if root == nil {
		return 0
	}

	leftMaxDepth := maxDepth(root.Left)
	rightMaxDepth := maxDepth(root.Right)

	return int(math.Max(float64(leftMaxDepth), float64(rightMaxDepth))) + 1
}

// invertTree 翻转二叉树
func invertTree(root *types.TreeNode) *types.TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}
