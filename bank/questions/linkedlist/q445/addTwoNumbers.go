package main

import "algorithm-golang/types"

/**
*
* 【445. 两数相加 II】
🌀https://leetcode.cn/problems/add-two-numbers-ii/
*
* @author  sun
* @date 2022/11/10 15:09
*/

// 先翻转链表，再计算两个链表的和，再将链表翻转作为结果
func addTwoNumbers(l1 *types.ListNode, l2 *types.ListNode) *types.ListNode {
	r1, r2 := reverse(l1), reverse(l2)
	var head *types.ListNode
	var tail *types.ListNode
	carry := 0
	for r1 != nil || r2 != nil {
		n1, n2 := 0, 0
		if r1 != nil {
			n1 = r1.Val
			r1 = r1.Next
		}
		if r2 != nil {
			n2 = r2.Val
			r2 = r2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &types.ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &types.ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &types.ListNode{Val: carry}
	}

	return reverse(head)
}

func reverse(head *types.ListNode) *types.ListNode {
	var record []int
	for head != nil {
		record = append(record, head.Val)
		head = head.Next
	}
	ans := &types.ListNode{Val: 0}
	p := ans
	for i := len(record) - 1; i >= 0; i-- {
		p.Next = &types.ListNode{Val: record[i]}
		p = p.Next
	}
	return ans.Next
}
