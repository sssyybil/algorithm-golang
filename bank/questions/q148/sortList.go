package main

import (
	"algorithm-golang/types"
	"sort"
)

func sortList(head *types.ListNode) *types.ListNode {
	p := head
	var nums []int
	for p != nil {
		nums = append(nums, p.Val)
		p = p.Next
	}
	sort.Ints(nums)
	p = head
	for _, v := range nums {
		p.Val = v
		p = p.Next
	}
	return head
}
