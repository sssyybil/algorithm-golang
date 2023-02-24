package main

import (
	"algorithm-golang/types"
	"fmt"
	"math"
)

/**
 * 【104. 二叉树的最大深度】🐱https://leetcode.cn/problems/maximum-depth-of-binary-tree/
 */

func maxDepth(root *types.TreeNode) int {
	if root == nil {
		return 0
	}

	leftMaxDepth := maxDepth(root.Left)
	rightMaxDepth := maxDepth(root.Right)

	return int(math.Max(float64(leftMaxDepth), float64(rightMaxDepth))) + 1
}

func main() {
	root := types.TreeNode{Val: 3}
	root.Left = &types.TreeNode{Val: 20}
	root.Right = &types.TreeNode{Val: 7}
	root.Right.Left = &types.TreeNode{Val: 9}
	root.Right.Right = &types.TreeNode{Val: 15}

	fmt.Println(maxDepth(&root))

}
