package main

import (
	"algorithm-golang/types"
	"algorithm-golang/utils"
	"fmt"
)

/**
*
* 【328. 奇偶链表】
🌀https://leetcode.cn/problems/odd-even-linked-list/description/
*
* @author  sun
* @date 2022/11/9 10:11
*/

func oddEvenListA(head *types.ListNode) *types.ListNode {
	odd := types.ListNode{}
	even := types.ListNode{}
	p, o, e := head, &odd, &even
	count := 1
	for p != nil {
		if count%2 != 0 {
			o.Next = &types.ListNode{Val: p.Val}
			o = o.Next
		} else {
			e.Next = &types.ListNode{Val: p.Val}
			e = e.Next
		}
		count++
		p = p.Next
	}
	o.Next = even.Next
	return odd.Next
}

func oddEvenListB(head *types.ListNode) *types.ListNode {
	if head == nil {
		return head
	}
	// 偶数节点的头节点
	evenHead := head.Next
	odd, even := head, evenHead
	// 当 even 为空节点或 even.Next 为空节点时，表示全部节点已经分离完毕，此时 odd 指向链表的最后一个节点
	// 为什么以偶数节点为空或偶数节点的下一个节点为空作为结束循环的条件呢❓
	for even != nil && even.Next != nil {
		// 更新奇数节点
		odd.Next = even.Next
		odd = odd.Next
		// 更新偶数节点
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}

func main() {
	utils.PrintLinkedList(oddEvenListA(utils.GenerateLinkedList([]int{1, 2, 3, 4, 5})))       // 1,3,5,2,4
	utils.PrintLinkedList(oddEvenListA(utils.GenerateLinkedList([]int{2, 1, 3, 5, 6, 4, 7}))) // 2,3,6,7,1,5,4
	fmt.Println()
	utils.PrintLinkedList(oddEvenListB(utils.GenerateLinkedList([]int{1, 2, 3, 4, 5})))       // 1,3,5,2,4
	utils.PrintLinkedList(oddEvenListB(utils.GenerateLinkedList([]int{2, 1, 3, 5, 6, 4, 7}))) // 2,3,6,7,1,5,4
}
