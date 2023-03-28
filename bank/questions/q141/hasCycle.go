package main

import "algorithm-golang/types"

/**
 * 【141. 环形链表】🐱https://leetcode.cn/problems/linked-list-cycle/description/?envType=study-plan&id=prepare-for-development&plan=development-job-hunting&plan_progress=xxd3p3rc
 */

func hasCycle(head *types.ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
