package binarytree

import (
	"algorithm-golang/types"
	"algorithm-golang/utils"
	"fmt"
	"testing"
)

func TestContain(t *testing.T) {
	node := types.TreeNode{Val: 1}
	node.Left = &types.TreeNode{Val: 2}
	node.Right = &types.TreeNode{Val: 3}
	node.Left.Left = &types.TreeNode{Val: 4}
	node.Left.Right = &types.TreeNode{Val: 5}
	node.Right.Left = &types.TreeNode{Val: 6}

	fmt.Println(contain(&node, 3))
}

func TestMinDepth(t *testing.T) {
	root := types.TreeNode{Val: 1}
	root.Left = &types.TreeNode{Val: 2}
	root.Right = &types.TreeNode{Val: 3}
	root.Left.Left = &types.TreeNode{Val: 4}
	root.Left.Right = &types.TreeNode{Val: 5}

	min := minDepth(&root)
	fmt.Println(min)
}

func TestInvertTree(t *testing.T) {
	root := types.TreeNode{Val: 4}
	root.Left = &types.TreeNode{Val: 2}
	root.Right = &types.TreeNode{Val: 7}
	root.Left.Left = &types.TreeNode{Val: 1}
	root.Left.Right = &types.TreeNode{Val: 3}
	root.Right.Left = &types.TreeNode{Val: 6}
	root.Right.Right = &types.TreeNode{Val: 9}
	fmt.Println(utils.LevelPrintBinaryTree(&root))

	ans := invertTree(&root)
	fmt.Println(utils.LevelPrintBinaryTree(ans))
}
