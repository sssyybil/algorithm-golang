package main

import (
	"algorithm-golang/types"
	"fmt"
)

/**
 * 【404. 左叶子之和】🐱https://leetcode.cn/problems/sum-of-left-leaves/
 */

// TODO Repeat

func sumOfLeftLeaves(root *types.TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root)
}

func isLeaveNode(node *types.TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

func dfs(node *types.TreeNode) (ans int) {
	if node.Left != nil {
		if isLeaveNode(node.Left) {
			ans += node.Left.Val
		} else {
			ans += dfs(node.Left)
		}
	}
	if node.Right != nil && !isLeaveNode(node.Right) {
		ans += dfs(node.Right)
	}
	return
}

func main() {
	root := types.TreeNode{Val: 3}
	root.Left = &types.TreeNode{Val: 9}
	root.Right = &types.TreeNode{Val: 20}
	root.Right.Left = &types.TreeNode{Val: 15}
	root.Right.Right = &types.TreeNode{Val: 7}

	fmt.Println(sumOfLeftLeaves(&root))
}
