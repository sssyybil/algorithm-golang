package main

import "fmt"

/**
 * 【876. 链表的中间结点】🐱https://leetcode.cn/problems/middle-of-the-linked-list/
 */

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// ✨✨【方法一：两次遍历】两次遍历，首先算出链表的长度，然后根据链表长度/2，即中间节点的下标，再次遍历链表，返回中间节点
func middleNode(head *ListNode) *ListNode {
	p, n := head, 0
	for p != nil {
		n++
		p = p.Next
	}
	for i := 0; i < n/2; i++ {
		head = head.Next
	}
	return head
}

// ✨【方法二：使用额外空间——数组】首次遍历单链表，计算链表的长度，并将链表中的各个节点存入数组中，最后返回数组下标为 n/2 的值即可
func middleNodeA(head *ListNode) *ListNode {
	var array []*ListNode
	p, n := head, 0
	for p != nil {
		array = append(array, p)
		n++
		p = p.Next
	}
	return array[n/2]
}

// ✨✨✨【方法二：快慢指针】慢指针每次走一步，快指针每次走两步，当快指针到达链表末尾的时候，慢指针必然位于中间（🌰 原因不明😭）
func middleNodeB(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func printListNode(node *ListNode) {
	p := node
	for p != nil {
		fmt.Printf("%v ", p.Val)
		p = p.Next
	}
	fmt.Println()
}

func main() {

	head := ListNode{Val: 1}
	p := &head
	for i := 1; i < 10; i++ {
		p.Next = &ListNode{Val: i}
		p = p.Next
	}

	printListNode(middleNode(&head))
	printListNode(middleNodeA(&head))
	printListNode(middleNodeB(&head))

}
