package main

import (
	"algorithm-golang/types"
)

// 【147. 对链表进行插入排序】👧🏻https://leetcode.cn/problems/insertion-sort-list/description/

// TODO repeat

func insertionSortList(head *types.ListNode) *types.ListNode {
	if head == nil {
		return head
	}
	dummyHead := &types.ListNode{Next: head}
	lastSorted, cur := head, head.Next
	for cur != nil {
		if lastSorted.Val <= cur.Val {
			lastSorted = lastSorted.Next
		} else {
			prev := dummyHead
			for prev.Next.Val <= cur.Val {
				prev = prev.Next
			}
			lastSorted.Next = cur.Next
			cur.Next = prev.Next
			prev.Next = cur
		}
		cur = lastSorted.Next
	}
	return dummyHead.Next
}
