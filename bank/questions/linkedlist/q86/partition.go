package main

import (
	"algorithm-golang/types"
	"algorithm-golang/utils"
	"fmt"
)

/**
*
* 【86. 分隔链表】
🌀https://leetcode.cn/problems/partition-list/
*
* @author  sun
* @date 2022/11/8 11:36
*/

func partitionA(head *types.ListNode, x int) *types.ListNode {
	var smaller []int
	var bigger []int
	p := head
	for p != nil {
		if p.Val < x {
			smaller = append(smaller, p.Val)
		} else {
			bigger = append(bigger, p.Val)
		}
		p = p.Next
	}

	p = head
	update := func(nums []int) *types.ListNode {
		for _, v := range nums {
			p.Val = v
			p = p.Next
		}
		return p
	}

	update(smaller)
	update(bigger)

	return head
}

func partitionB(head *types.ListNode, x int) *types.ListNode {
	smaller := types.ListNode{}
	larger := types.ListNode{}
	p, s, l := head, &smaller, &larger
	for p != nil {
		if p.Val < x {
			s.Next = &types.ListNode{Val: p.Val}
			s = s.Next
		} else {
			l.Next = &types.ListNode{Val: p.Val}
			l = l.Next
		}
		p = p.Next
	}
	s.Next = larger.Next
	return smaller.Next
}

func main() {
	utils.PrintLinkedList(partitionA(utils.GenerateLinkedList([]int{1, 4, 3, 2, 5, 2}), 3)) // 1,2,2,4,3,5
	utils.PrintLinkedList(partitionA(utils.GenerateLinkedList([]int{2, 1}), 2))             // 1,2

	fmt.Printf("-------------->\n")

	utils.PrintLinkedList(partitionB(utils.GenerateLinkedList([]int{1, 4, 3, 2, 5, 2}), 3)) // 1,2,2,4,3,5
	utils.PrintLinkedList(partitionB(utils.GenerateLinkedList([]int{2, 1}), 2))             // 1,2
}
