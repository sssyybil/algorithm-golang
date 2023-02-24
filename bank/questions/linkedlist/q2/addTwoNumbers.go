package main

import (
	"algorithm-golang/types"
	"strconv"
	"strings"
)

/**
*
* 【2. 两数相加】
🌀https://leetcode.cn/problems/add-two-numbers/
*
* @author  sun
* @date 2022/11/10 10:13
*/

// ⚠️ 理解错提议了，代码非常丑陋，留着示众
//
// ❗️测试失败，卡在了测试用例：
// l1=1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1
// l2=5,6,4
// 原因：l1 数值过大，使用 int64 存储会报错：panic: strconv.ParseInt: parsing "1000000000000000000000000000001": value out of range
func addTwoNumbersA(l1 *types.ListNode, l2 *types.ListNode) *types.ListNode {
	s1, s2 := "", ""
	p1, p2 := l1, l2
	for p1 != nil {
		s1 = strconv.Itoa(p1.Val) + s1
		p1 = p1.Next
	}
	for p2 != nil {
		s2 = strconv.Itoa(p2.Val) + s2
		p2 = p2.Next
	}

	n1, err := strconv.ParseInt(s1, 10, 64)
	if err != nil {
		panic(err)
	}
	n2, err := strconv.ParseInt(s2, 10, 64)
	if err != nil {
		panic(err)
	}

	sum := strings.Split(strconv.FormatInt(n1+n2, 10), "")
	ans := types.ListNode{}
	p := &ans
	for i := len(sum) - 1; i >= 0; i-- {
		val, err := strconv.Atoi(sum[i])
		if err != nil {
			panic(err)
		}
		p.Next = &types.ListNode{Val: val}
		p = p.Next
	}

	return ans.Next
}

func addTwoNumbersB(l1 *types.ListNode, l2 *types.ListNode) *types.ListNode {
	var tail *types.ListNode
	var head *types.ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &types.ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &types.ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &types.ListNode{Val: carry}
	}
	return head
}
