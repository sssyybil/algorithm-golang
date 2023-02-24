package main

import "fmt"

/**
 * 【19. 删除链表的倒数第 N 个结点】🐱https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// removeNthFromEnd -> 双指针，为删除 slow 指针指向的下一个节点，设置 dummy 节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 哑节点
	dummy := &ListNode{Val: 0, Next: head}
	slow, fast := dummy, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for ; fast != nil; fast = fast.Next {
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

func createListNode(nums []int) *ListNode {
	head := ListNode{}
	p := &head
	for _, v := range nums {
		p.Next = &ListNode{Val: v}
		p = p.Next
	}
	return head.Next
}

func printListNode(head *ListNode) {
	for head != nil {
		fmt.Printf("%v ", head.Val)
		head = head.Next
	}
}

func main() {
	head := createListNode([]int{1, 2, 3, 4, 5})
	res := removeNthFromEnd(head, 5)
	printListNode(res)
}
