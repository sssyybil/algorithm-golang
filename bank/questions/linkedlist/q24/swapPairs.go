package main

import "algorithm-golang/types"

//【24. 两两交换链表中的节点】 🌀https://leetcode.cn/problems/swap-nodes-in-pairs/description/

func swapPairs(head *types.ListNode) *types.ListNode {
	p := head
	for p != nil && p.Next != nil {
		next := p.Next
		p.Val, next.Val = next.Val, p.Val
		p = next.Next
	}
	return head
}
