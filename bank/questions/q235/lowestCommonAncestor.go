package main

import "fmt"

/**
 * 【235. 二叉搜索树的最近公共祖先】🐱https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pathP := getPath(root, p)
	pathQ := getPath(root, q)

	var ans *TreeNode
	for i := 0; i < len(pathP) && i < len(pathQ) && pathP[i] == pathQ[i]; i++ {
		ans = pathP[i]
	}
	return ans
}

func getPath(root, target *TreeNode) (path []*TreeNode) {
	node := root
	for node.Val != target.Val {
		path = append(path, node)
		if target.Val < node.Val {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	path = append(path, node)
	return
}

// todo 构建二叉树
func packageNode(array []int) *TreeNode {
	if array == nil || len(array) == 0 {
		return nil
	}
	return nil
}

func main() {
	root := TreeNode{Val: 6}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 8}
	root.Left.Left = &TreeNode{Val: 0}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 7}
	root.Right.Right = &TreeNode{Val: 9}
	root.Left.Right.Left = &TreeNode{Val: 3}
	root.Left.Right.Right = &TreeNode{Val: 5}

	pathP := getPath(&root, &TreeNode{Val: 2})
	pathQ := getPath(&root, &TreeNode{Val: 4})

	for i, v := range pathP {
		fmt.Printf("%v - %d\n", i, v.Val)
	}

	fmt.Println()

	for i, v := range pathQ {
		fmt.Printf("%v - %d\n", i, v.Val)
	}

	fmt.Println()

	res := lowestCommonAncestor(&root, &TreeNode{Val: 2}, &TreeNode{Val: 4})
	fmt.Println(res.Val)
}
