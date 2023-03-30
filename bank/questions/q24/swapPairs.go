package main

import "algorithm-golang/types"

/**
 * 【24. 两两交换链表中的节点】
 * 🪢https://leetcode.cn/problems/swap-nodes-in-pairs/description/?envType=study-plan&id=prepare-for-development&plan=development-job-hunting&plan_progress=xxd3p3rc
 * @2023-03-28 18:12
 */

func swapPairs(head *types.ListNode) *types.ListNode {
	if head == nil {
		return head
	}
	dummyNode := &types.ListNode{Next: head}
	p := dummyNode
	for p.Next != nil && p.Next.Next != nil {
		p1 := p.Next
		p2 := p1.Next

		p.Next = p2
		p1.Next = p2.Next
		p2.Next = p1

		p = p1
	}

	return dummyNode.Next
}
