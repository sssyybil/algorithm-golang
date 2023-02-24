package main

import "algorithm-golang/types"

// 使用【插入排序】，提交后超时
func sortListWithInsert(head *types.ListNode) *types.ListNode {
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

func sortListWithMerge(head *types.ListNode) *types.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	tmp := slow.Next
	slow.Next = nil
	left := sortListWithMerge(head)
	right := sortListWithMerge(tmp)
	ans := &types.ListNode{}
	p := ans
	for left != nil || right != nil {
		if left == nil {
			p.Next = right
		} else if right == nil {
			p.Next = left
		} else if left.Val < right.Val {
			p.Next = left
			left = left.Next
		} else {
			p.Next = right
			right = right.Next
		}
		p = p.Next
	}
	return ans.Next
}
