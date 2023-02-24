package main

import (
	"algorithm-golang/types"
	"fmt"
)

/**
* 【100. 相同的树】🐱https://leetcode.cn/problems/same-tree/
 */

func isSameTree(p *types.TreeNode, q *types.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (p == nil && q != nil) || (p != nil && q == nil) || (p.Val != q.Val) {
		return false
	}
	if !isSameTree(p.Left, q.Left) || !isSameTree(p.Right, q.Right) {
		return false
	}
	return true
}

func main() {
	p := types.TreeNode{Val: 1}
	p.Left = &types.TreeNode{Val: 2}

	q := types.TreeNode{Val: 1}
	q.Right = &types.TreeNode{Val: 2}

	fmt.Println(isSameTree(&p, &q))
}
