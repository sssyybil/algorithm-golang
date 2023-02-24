package main

import (
	"algorithm-golang/types"
	"algorithm-golang/utils"
	"testing"
)

func Test_addTwoNumbers1(t *testing.T) {
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
			l1: utils.CreateLinkedList([]int{7, 2, 4, 3}),
			l2: utils.CreateLinkedList([]int{5, 6, 4}),
		}, utils.CreateLinkedList([]int{7, 8, 0, 7})},

		{"2", args{
			l1: utils.CreateLinkedList([]int{2, 4, 3}),
			l2: utils.CreateLinkedList([]int{5, 6, 4}),
		}, utils.CreateLinkedList([]int{8, 0, 7})},

		{"3", args{
			l1: utils.CreateLinkedList([]int{0}),
			l2: utils.CreateLinkedList([]int{0}),
		}, utils.CreateLinkedList([]int{0})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !utils.IsLinkedListEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", utils.ConvertLinkedListToSlice(got), utils.ConvertLinkedListToSlice(tt.want))
			}
		})
	}
}
