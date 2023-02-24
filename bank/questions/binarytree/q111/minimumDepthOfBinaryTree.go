package main

import (
	"algorithm-golang/types"
	"fmt"
	"math"
)

/**
 * 【111. 二叉树的最小深度】🐱https://leetcode.cn/problems/minimum-depth-of-binary-tree/
 */

func minDepth(root *types.TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return minDepth(root.Right) + 1
	} else if root.Right == nil {
		return minDepth(root.Left) + 1
	} else {
		leftMinDepth := minDepth(root.Left)
		rightMinDepth := minDepth(root.Right)
		return int(math.Min(float64(leftMinDepth), float64(rightMinDepth))) + 1
	}
}

func main() {
	root1 := types.TreeNode{Val: 3}
	root1.Left = &types.TreeNode{Val: 9}
	root1.Right = &types.TreeNode{Val: 20}
	root1.Right.Left = &types.TreeNode{Val: 15}
	root1.Right.Right = &types.TreeNode{Val: 7}

	root2 := types.TreeNode{Val: 2}
	root2.Right = &types.TreeNode{Val: 3}
	root2.Right.Right = &types.TreeNode{Val: 4}
	root2.Right.Right.Right = &types.TreeNode{Val: 5}
	root2.Right.Right.Right.Right = &types.TreeNode{Val: 6}

	fmt.Println(
		minDepth(&root1),
		minDepth(&root2),
	)

}
