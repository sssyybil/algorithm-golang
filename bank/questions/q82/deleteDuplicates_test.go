package main

import (
	"algorithm-golang/types"
	"algorithm-golang/utils"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *types.ListNode
	}
	tests := []struct {
		name string
		args args
		want *types.ListNode
	}{
		// TODO: Add test cases.
		{"1", args{head: utils.CreateLinkedList([]int{1, 1, 1, 2, 3})}, utils.CreateLinkedList([]int{2, 3})},
		{"2", args{head: utils.CreateLinkedList([]int{1, 2, 3, 3, 4, 5})}, utils.CreateLinkedList([]int{1, 2, 4, 5})},
		{"3", args{head: utils.CreateLinkedList([]int{2, 3, 4, 4, 4})}, utils.CreateLinkedList([]int{2, 3})},
		{"4", args{head: utils.CreateLinkedList([]int{3, 4, 5, 6, 7})}, utils.CreateLinkedList([]int{3, 4, 5, 6, 7})},
		{"5", args{head: utils.CreateLinkedList([]int{1, 2, 3, 3, 4, 4, 5})}, utils.CreateLinkedList([]int{1, 2, 5})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !utils.IsLinkedListEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", utils.ConvertLinkedListToSlice(got), utils.ConvertLinkedListToSlice(tt.want))
			}
		})
	}
}
