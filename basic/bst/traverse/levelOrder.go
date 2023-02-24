package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 广度优先遍历
func levelOrder(root *TreeNode) {
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		var pop *TreeNode
		pop = queue[0]
		// 上述两行代码等同于：pop := *queue[0]
		queue = queue[1:]
		fmt.Printf("%d, %v\n", pop.Val, pop)
		if pop.Left != nil {
			queue = append(queue, pop.Left)
		}
		if pop.Right != nil {
			queue = append(queue, pop.Right)
		}
	}
}

func main() {
	root := TreeNode{Val: 28}
	root.Left = &TreeNode{Val: 16}
	root.Left.Left = &TreeNode{Val: 13}
	root.Left.Right = &TreeNode{Val: 22}
	root.Right = &TreeNode{Val: 30}
	root.Right.Left = &TreeNode{Val: 29}
	root.Right.Right = &TreeNode{Val: 42}

	levelOrder(&root)
}
