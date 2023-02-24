package main

import (
	"algorithm-golang/types"
	"algorithm-golang/utils"
	"fmt"
)

/**
*
* 【92. 反转链表 II】
🌀https://leetcode.cn/problems/reverse-linked-list-ii/description/
*
* @author  sun
* @date 2022/11/7 10:26
*/

// 使用 栈 作为额外的数据空间，⚠️通常情况下（比如面试），在做这类题目时是不允许修改节点值的，只能修改节点的指向
func reverseBetweenA(head *types.ListNode, left int, right int) *types.ListNode {
	var stack []int
	count, p := 1, head
	for p != nil {
		if count >= left && count <= right {
			stack = append(stack, p.Val)
		}
		p = p.Next
		count++
	}
	fmt.Printf("stack: %v\n", stack)
	count, p = 1, head
	for p != nil {
		if count >= left && count <= right {
			n := len(stack) - 1
			p.Val = stack[n]
			stack = stack[:n]
		}
		p = p.Next
		count++
	}
	return head
}

func reverseLinkedList(head *types.ListNode) {
	var pre *types.ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
}

func reverseBetweenB(head *types.ListNode, left, right int) *types.ListNode {
	// 因为头节点有可能发生变化，使用虚拟头节点可以避免复杂的分类讨论
	dummyNode := &types.ListNode{Val: -1}
	dummyNode.Next = head

	pre := dummyNode
	// 第 1 步：从虚拟头节点走 left - 1 步，来到 left 节点的前一个节点
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	// 第 2 步：从 pre 再走 right - left + 1 步，来到 right 节点
	rightNode := pre
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}

	// 第 3 步：切断出一个子链表（截取链表）
	leftNode := pre.Next
	curr := rightNode.Next

	// 注意：切断链接
	pre.Next = nil
	rightNode.Next = nil

	// 第 4 步：同第 206 题，反转链表的子区间
	reverseLinkedList(leftNode)

	// 第 5 步：接回到原来的链表中
	pre.Next = rightNode
	leftNode.Next = curr
	return dummyNode.Next
}

func main() {
	case1 := utils.CreateLinkedList([]int{1, 2, 3, 4, 5})
	case2 := utils.CreateLinkedList([]int{5})

	//utils.PrintLinkedList(reverseBetweenA(case1, 2, 4))

	utils.PrintLinkedList(reverseBetweenB(case1, 2, 4))
	utils.PrintLinkedList(reverseBetweenB(case2, 1, 1))

}
