package main

import (
	"algorithm-golang/types"
	"algorithm-golang/utils"
	"testing"
)

func Test_addTwoNumbersB(t *testing.T) {
	type args struct {
		l1 *types.ListNode
		l2 *types.ListNode
	}
	tests := []struct {
		name string
		args args
		want *types.ListNode
	}{
		{"1", args{
			l1: utils.GenerateLinkedList([]int{2, 4, 3}),
			l2: utils.GenerateLinkedList([]int{5, 6, 4}),
		}, utils.GenerateLinkedList([]int{7, 0, 8})},

		{"2", args{
			l1: utils.GenerateLinkedList([]int{7, 4, 5}),
			l2: utils.GenerateLinkedList([]int{8, 9, 2}),
		}, utils.GenerateLinkedList([]int{5, 4, 8})},

		{"3", args{
			l1: utils.GenerateLinkedList([]int{0}),
			l2: utils.GenerateLinkedList([]int{0}),
		}, utils.GenerateLinkedList([]int{0})},

		{"4", args{
			l1: utils.GenerateLinkedList([]int{9, 9, 9, 9, 9, 9, 9}),
			l2: utils.GenerateLinkedList([]int{9, 9, 9, 9}),
		}, utils.GenerateLinkedList([]int{8, 9, 9, 9, 0, 0, 0, 1})},

		{"5", args{
			l1: utils.GenerateLinkedList([]int{8, 3, 9}),
			l2: utils.GenerateLinkedList([]int{2, 5, 1}),
		}, utils.GenerateLinkedList([]int{0, 9, 0, 1})},

		{"6", args{
			l1: utils.GenerateLinkedList([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}),
			l2: utils.GenerateLinkedList([]int{5, 6, 4}),
		}, utils.GenerateLinkedList([]int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})},
	}
	// 测试逻辑
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbersB(tt.args.l1, tt.args.l2); !utils.IsLinkedListEqual(got, tt.want) {
				t.Errorf("addTwoNumbersB() = %v, want %v", utils.ConvertLinkedListToSlice(got), utils.ConvertLinkedListToSlice(tt.want))
			}
		})
	}
}
