package main

import (
	"algorithm-golang/types"
	"algorithm-golang/utils"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *types.ListNode
		list2 *types.ListNode
	}
	tests := []struct {
		name string
		args args
		want *types.ListNode
	}{
		// TODO: Add test cases.
		{"1", args{list1: utils.GenerateLinkedList([]int{1, 2, 4}), list2: utils.GenerateLinkedList([]int{1, 3, 4})}, utils.GenerateLinkedList([]int{1, 1, 2, 3, 4, 4})},
		{"2", args{list1: utils.GenerateLinkedList([]int{}), list2: utils.GenerateLinkedList([]int{})}, utils.GenerateLinkedList([]int{})},
		{"3", args{list1: utils.GenerateLinkedList([]int{1, 2, 3, 4, 5}), list2: utils.GenerateLinkedList([]int{})}, utils.GenerateLinkedList([]int{1, 2, 3, 4, 5})},
		{"4", args{list1: utils.GenerateLinkedList([]int{2, 2, 2}), list2: utils.GenerateLinkedList([]int{3, 3, 3})}, utils.GenerateLinkedList([]int{2, 2, 2, 3, 3, 3})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.list1, tt.args.list2); !utils.IsLinkedListEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", utils.ConvertLinkedListToSlice(got), utils.ConvertLinkedListToSlice(tt.want))
			}
		})
	}
}
