package main

import "algorithm-golang/types"

//【203. 移除链表元素】 🌀https://leetcode.cn/problems/remove-linked-list-elements/description/

func removeElements(head *types.ListNode, val int) *types.ListNode {
	dummyNode := types.ListNode{
		Val:  0,
		Next: head,
	}
	cur := &dummyNode
	for cur != nil && cur.Next != nil {
		for cur.Next != nil && cur.Next.Val == val {
			delNode := cur.Next
			cur.Next = delNode.Next
		}
		cur = cur.Next
	}
	return dummyNode.Next
}
