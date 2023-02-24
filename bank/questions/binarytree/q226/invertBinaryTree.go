package main

import (
	"algorithm-golang/types"
	"algorithm-golang/utils"
	"fmt"
)

/**
 * 【226. 翻转二叉树】🐱https://leetcode.cn/problems/invert-binary-tree/
 */

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

func main() {
	root := types.TreeNode{Val: 4}
	root.Left = &types.TreeNode{Val: 2}
	root.Right = &types.TreeNode{Val: 7}
	root.Left.Left = &types.TreeNode{Val: 1}
	root.Left.Right = &types.TreeNode{Val: 3}
	root.Right.Left = &types.TreeNode{Val: 6}
	root.Right.Right = &types.TreeNode{Val: 9}

	ans := invertTree(&root)

	fmt.Println(utils.LevelPrintBinaryTree(ans))
}
