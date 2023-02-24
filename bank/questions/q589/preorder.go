package main

/**
 * 【589. N 叉树的前序遍历】🐱https://leetcode.cn/problems/n-ary-tree-preorder-traversal/
 */

type Node struct {
	Val      int
	Children []*Node
}

// 递归
func preorderDfs(root *Node) []int {
	var ans []int
	var dfs func(*Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		for _, ch := range node.Children {
			dfs(ch)
		}
	}
	dfs(root)
	return ans
}

func main() {

}
