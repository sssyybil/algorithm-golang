package main

import (
	"algorithm-golang/types"
	"fmt"
)

/**
 * 【222. 完全二叉树的节点个数】🐱https://leetcode.cn/problems/count-complete-tree-nodes/submissions/
 *
 *  * 完全二叉树：除了最后一层，所有层的节点数达到最大，与此同时，最后一层的所有节点都在最左侧。（堆使用了完全二叉树）
 *  * 满二叉树：所有层的节点数达到最大
 *  * 平衡二叉树：每一个节点的左右子树的高度差不超过一
 */

// TODO 使用递归求解未利用完全二叉树的特性，若利用其特性的话该如何求解？

func countNodes(root *types.TreeNode) int {
	if root == nil {
		return 0
	}
	left := countNodes(root.Left)
	right := countNodes(root.Right)
	return left + right + 1
}

func main() {
	root := types.TreeNode{Val: 1}
	root.Left = &types.TreeNode{Val: 2}
	root.Right = &types.TreeNode{Val: 3}
	root.Left.Left = &types.TreeNode{Val: 4}
	root.Left.Right = &types.TreeNode{Val: 5}
	root.Right.Left = &types.TreeNode{Val: 6}
	fmt.Println(countNodes(&root))

}
