package main

import (
	"algorithm-golang/types"
	"algorithm-golang/utils"
)

/**
 * 【143. 重排链表】🐱https://leetcode.cn/problems/reorder-list/description/
 */

// reorderListWithArray 使用数组（线性表）存储链表中的元素，然后利用可以数组可以通过下标访问的特性，重组链表
func reorderListWithArray(head *types.ListNode) {
	var record []int
	p := head
	for p != nil {
		record = append(record, p.Val)
		p = p.Next
	}

	p, l, r := head, 0, len(record)-1
	// 当链表长度为奇数时，需判断 p.Next != nil；当链表长度为偶数时，需判断 p != nil
	for l < r && p != nil && p.Next != nil {
		p.Val = record[l]
		p.Next.Val = record[r]
		p = p.Next.Next
		l++
		r--
	}
	if l == r {
		p.Val = record[l]
	}
	utils.PrintLinkedList(head)
}

// reorderList 寻找链表的中间节点 + 将链表的右半部分反转 + 合并链表
func reorderList(head *types.ListNode) {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	first, second := head, reverse(slow.Next)
	slow.Next = nil
	merge(first, second)
}

func reverse(head *types.ListNode) *types.ListNode {
	var dummyNode *types.ListNode
	p, cur := dummyNode, head
	for cur != nil {
		next := cur.Next
		cur.Next = p
		p = cur
		cur = next
	}
	return p
}

// merge h1=a1->a2->a3->null，h2=b1->b2->b3->null，合并 h1 和 h2，结果应为：a1->b1->a2->b2->a3->b3->null
func merge(h1, h2 *types.ListNode) {
	var h1Tmp, h2Tmp *types.ListNode
	for h1 != nil && h2 != nil {
		h1Tmp = h1.Next
		h2Tmp = h2.Next

		h1.Next = h2
		h1 = h1Tmp

		h2.Next = h1
		h2 = h2Tmp
	}
}
