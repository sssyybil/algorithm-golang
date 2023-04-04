package utils

import (
	"algorithm-golang/types"
)

//func PackageBinaryTreeNode(nums []int) {
//	root := types.TreeNode{}
//	p := root
//	for _, v := range nums {
//
//	}
//}

// LevelPrintBinaryTree 二叉树的层序遍历（🎩https://leetcode.cn/problems/binary-tree-level-order-traversal/）
func LevelPrintBinaryTree(root *types.TreeNode) [][]int {
	var ans [][]int
	queue := []*types.TreeNode{root}
	for len(queue) > 0 {

		n := len(queue)
		var level []int

		for i := n; i > 0; i-- {
			node := queue[:1][0]
			queue = queue[1:]
			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, level)
	}
	return ans
}
