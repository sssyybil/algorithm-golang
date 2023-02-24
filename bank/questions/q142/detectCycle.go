package main

import "fmt"

/**
 * 【142. 环形链表 II】🐱https://leetcode.cn/problems/linked-list-cycle-ii/
 */

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		if fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}

func main() {

	head := ListNode{Val: 3}
	node1 := ListNode{Val: 2}
	node2 := ListNode{Val: 0}
	node3 := ListNode{Val: -4}
	head.Next = &node1
	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node1

	res := detectCycle(&head)
	fmt.Printf("%v", res)
}
