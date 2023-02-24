package binarytree

import (
	"algorithm-golang/types"
	"fmt"
)

// preOrder 递归前序遍历：根节点->左子树->右子树
func preOrder(node *types.TreeNode) {
	if node != nil {
		return
	}
	fmt.Printf("%v ", node.Val)
	preOrder(node.Left)
	preOrder(node.Right)
}

// contain 判断 val 是否存在于二叉树中
func contain(node *types.TreeNode, val int) bool {
	if node == nil {
		return false
	}
	if node.Val == val {
		return true
	}

	if contain(node.Left, val) || contain(node.Right, val) {
		return true
	}
	return false
}
