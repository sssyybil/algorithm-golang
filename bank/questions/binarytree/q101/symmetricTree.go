package main

import (
	"algorithm-golang/types"
	"fmt"
)

/**
 * 【101. 对称二叉树】🐱https://leetcode.cn/problems/symmetric-tree/
 */

func isSymmetric(root *types.TreeNode) bool {
	return check(root.Left, root.Right)
}

func check(p, q *types.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}

func main() {
	root := types.TreeNode{Val: 1}
	root.Left = &types.TreeNode{Val: 2}
	root.Right = &types.TreeNode{Val: 2}
	root.Left.Left = &types.TreeNode{Val: 3}
	root.Left.Right = &types.TreeNode{Val: 4}
	root.Right.Left = &types.TreeNode{Val: 4}
	root.Right.Right = &types.TreeNode{Val: 3}
	fmt.Println(isSymmetric(&root))

	root2 := types.TreeNode{Val: 1}
	root2.Left = &types.TreeNode{Val: 2}
	root2.Right = &types.TreeNode{Val: 2}
	root2.Left.Right = &types.TreeNode{Val: 3}
	root2.Right.Right = &types.TreeNode{Val: 3}
	fmt.Println(isSymmetric(&root2))

	root3 := types.TreeNode{Val: 1}
	root3.Left = &types.TreeNode{Val: 2}
	root3.Right = &types.TreeNode{Val: 2}
	root3.Left.Left = &types.TreeNode{Val: 2}
	root3.Right.Left = &types.TreeNode{Val: 2}
	fmt.Println(isSymmetric(&root3))

}
