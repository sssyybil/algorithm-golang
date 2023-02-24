package main

import (
	"fmt"
	"math"
)

/**
 * 【98. 验证二叉搜索树】🐱https://leetcode.cn/problems/validate-binary-search-tree/
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	fmt.Printf("root = %v, lower = %v, upper = %v\n", root, lower, upper)
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}

func main() {
	root1 := TreeNode{Val: 2}
	root1.Left = &TreeNode{Val: 1}
	root1.Right = &TreeNode{Val: 3}

	root2 := TreeNode{Val: 5}
	root2.Left = &TreeNode{Val: 1}
	root2.Right = &TreeNode{Val: 4}
	root2.Right.Left = &TreeNode{Val: 3}
	root2.Right.Right = &TreeNode{Val: 6}

	root3 := TreeNode{Val: 6}
	root3.Left = &TreeNode{Val: 5}
	root3.Right = &TreeNode{Val: 7}
	root3.Left.Left = &TreeNode{Val: 4}
	root3.Left.Right = &TreeNode{Val: 9}
	root3.Right.Left = &TreeNode{Val: 3}
	root3.Right.Right = &TreeNode{Val: 8}

	fmt.Println(
		//isValidBST(&root1),
		//isValidBST(&root2),
		isValidBST(&root3),
	)
}
