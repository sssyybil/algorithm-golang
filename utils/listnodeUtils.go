package utils

import (
	"algorithm-golang/types"
	"fmt"
)

// GenerateLinkedList ==> 根据给定数组生成单链表
func GenerateLinkedList(nums []int) *types.ListNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	dummyHead := &types.ListNode{}
	p := dummyHead
	for _, v := range nums {
		p.Next = &types.ListNode{Val: v}
		p = p.Next
	}
	return dummyHead.Next
}

// GenerateCycleLinkedList 根据给定数组生成循环单链表
func GenerateCycleLinkedList(nums []int) *types.ListNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	dummyHead := &types.ListNode{}
	p := dummyHead
	for _, v := range nums {
		p.Next = &types.ListNode{Val: v}
		p = p.Next
	}
	p.Next = dummyHead.Next
	return dummyHead.Next
}

// PrintLinkedList ==> 打印单链表
func PrintLinkedList(head *types.ListNode) {
	for head != nil {
		fmt.Printf("%v->", head.Val)
		head = head.Next
	}
	fmt.Printf("NULL\n")
}

// ConvertLinkedListToSlice ==> 将单链表转为切片存储
func ConvertLinkedListToSlice(head *types.ListNode) []int {
	var ans []int
	for head != nil {
		ans = append(ans, head.Val)
		head = head.Next
	}
	return ans
}

func IsLinkedListEqual(l1, l2 *types.ListNode) bool {
	for l1 != nil || l2 != nil {
		if (l1 != nil && l2 == nil) || (l1 == nil && l2 != nil) {
			return false
		}
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return true
}
