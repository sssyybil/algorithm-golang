package main

import (
	"algorithm-golang/types"
)

/**
 * 【82. 删除排序链表中的重复元素 II】🐱https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/description/
 */

func deleteDuplicates(head *types.ListNode) *types.ListNode {
	dummyNode := types.ListNode{Val: 0, Next: head}
	cur := &dummyNode
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummyNode.Next
}
