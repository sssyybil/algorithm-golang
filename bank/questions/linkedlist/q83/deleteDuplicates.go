package main

import (
	"algorithm-golang/types"
	"algorithm-golang/utils"
	"fmt"
)

/**
*
* 【83. 删除排序链表中的重复元素】
🌀https://leetcode.cn/problems/remove-duplicates-from-sorted-list/description/
*
* @author  sun
* @date 2022/11/8 10:25
*/

func deleteDuplicatesA(head *types.ListNode) *types.ListNode {
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

// 使用额外的存储空间 map 记录已经遍历过的节点值，判重使用
func deleteDuplicatesB(head *types.ListNode) *types.ListNode {
	record := make(map[int]bool)
	cur := head
	for cur != nil && cur.Next != nil {
		record[cur.Val] = true
		next := cur.Next
		if record[next.Val] {
			cur.Next = next.Next
		} else {
			cur = next
		}
	}
	return head
}

func main() {
	utils.PrintLinkedList(deleteDuplicatesA(utils.CreateLinkedList([]int{1, 1, 1})))                   // 1
	utils.PrintLinkedList(deleteDuplicatesA(utils.CreateLinkedList([]int{2, 2, 2, 2, 2, 3, 3, 3, 3}))) // 2,3
	utils.PrintLinkedList(deleteDuplicatesA(utils.CreateLinkedList([]int{1, 1, 2, 3, 3, 4, 5})))       // 1,2,3,4,5
	utils.PrintLinkedList(deleteDuplicatesA(utils.CreateLinkedList([]int{1, 1, 2})))                   // 1,2
	utils.PrintLinkedList(deleteDuplicatesA(utils.CreateLinkedList([]int{1, 1, 2, 3, 3})))             // 1,2,3
	fmt.Println("------")
	utils.PrintLinkedList(deleteDuplicatesB(utils.CreateLinkedList([]int{1, 1, 1})))                   // 1
	utils.PrintLinkedList(deleteDuplicatesB(utils.CreateLinkedList([]int{2, 2, 2, 2, 2, 3, 3, 3, 3}))) // 2,3
	utils.PrintLinkedList(deleteDuplicatesB(utils.CreateLinkedList([]int{1, 1, 2, 3, 3, 4, 5})))       // 1,2,3,4,5
	utils.PrintLinkedList(deleteDuplicatesB(utils.CreateLinkedList([]int{1, 1, 2})))                   // 1,2
	utils.PrintLinkedList(deleteDuplicatesB(utils.CreateLinkedList([]int{1, 1, 2, 3, 3})))             // 1,2,3
}
