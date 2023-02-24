package main

import (
	"algorithm-golang/types"
	"fmt"
)

/**
 * 【112. 路径总和】🐱https://leetcode.cn/problems/path-sum/
 */

func hasPathSum(root *types.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	// 左右孩子节点均为空，当前节点才为叶子节点
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	if hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val) {
		return true
	}

	return false
}

func main() {
	root := types.TreeNode{Val: 5}
	root.Left = &types.TreeNode{Val: 4}
	root.Right = &types.TreeNode{Val: 8}
	root.Left.Left = &types.TreeNode{Val: 11}
	root.Right.Left = &types.TreeNode{Val: 13}
	root.Right.Right = &types.TreeNode{Val: 4}
	root.Left.Left.Left = &types.TreeNode{Val: 7}
	root.Left.Left.Right = &types.TreeNode{Val: 2}
	root.Right.Right.Right = &types.TreeNode{Val: 1}

	fmt.Println(hasPathSum(&root, 22))

}
