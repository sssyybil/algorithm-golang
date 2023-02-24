package main

import (
	"algorithm-golang/types"
	"algorithm-golang/utils"
)

/**
*
* 【206. 反转链表】
🌀https://leetcode.cn/problems/reverse-linked-list/description/
*
* @author  sun
* @date 2022/11/7 10:10
*/

func reverseList(head *types.ListNode) *types.ListNode {
	var pre *types.ListNode
	pre = nil
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func main() {
	listNode := utils.CreateLinkedList([]int{1, 2, 3, 4, 5})
	ans := reverseList(listNode)
	utils.PrintLinkedList(ans)
}
