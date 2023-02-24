package main

import (
	"algorithm-golang/types"
	"algorithm-golang/utils"
)

// 【面试题 02.01. 移除重复节点】🌀https://leetcode.cn/problems/remove-duplicate-node-lcci/description/?favorite=xb9lfcwi

func removeDuplicateNodes(head *types.ListNode) *types.ListNode {
	tmp := make(map[int]bool)
	pre, cur := &types.ListNode{Next: head}, head
	for cur != nil {
		if !tmp[cur.Val] {
			tmp[cur.Val] = true
			pre = pre.Next
			cur = cur.Next
		} else {
			next := cur.Next
			pre.Next = next
			cur = next
		}
	}
	return head
}

func main() {
	linkedList := utils.CreateLinkedList([]int{1, 1, 1, 1, 2})
	res := removeDuplicateNodes(linkedList)
	utils.PrintLinkedList(res)
}
