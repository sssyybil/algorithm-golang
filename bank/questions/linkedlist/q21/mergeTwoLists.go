package main

import "algorithm-golang/types"

/**
 * 【21. 合并两个有序链表】🐱https://leetcode.cn/problems/merge-two-sorted-lists/
 */

func mergeTwoLists(list1 *types.ListNode, list2 *types.ListNode) *types.ListNode {
	ans := types.ListNode{}
	p := &ans
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			p.Next = list1
			list1 = list1.Next
		} else {
			p.Next = list2
			list2 = list2.Next
		}
		p = p.Next
	}
	p.Next = list1
	if p.Next == nil {
		p.Next = list2
	}
	return ans.Next
}
