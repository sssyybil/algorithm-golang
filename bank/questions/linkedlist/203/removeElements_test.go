package main

import (
	"algorithm-golang/types"
	"algorithm-golang/utils"
	"testing"
)

func Test_removeElements(t *testing.T) {
	type args struct {
		head *types.ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *types.ListNode
	}{
		// TODO: Add test cases.
		{"removeElements", args{utils.GenerateLinkedList([]int{1, 2, 6, 3, 4, 5, 6}), 6}, utils.GenerateLinkedList([]int{1, 2, 3, 4, 5})},
		{"removeElements", args{utils.GenerateLinkedList([]int{1, 1, 1, 2, 2, 4, 5, 6}), 1}, utils.GenerateLinkedList([]int{2, 2, 4, 5, 6})},
		{"removeElements", args{utils.GenerateLinkedList([]int{3, 3, 3, 3}), 1}, utils.GenerateLinkedList([]int{3, 3, 3, 3})},
		{"removeElements", args{utils.GenerateLinkedList([]int{3, 3, 3, 3}), 3}, utils.GenerateLinkedList([]int{})},
		{"removeElements", args{utils.GenerateLinkedList([]int{}), 1}, utils.GenerateLinkedList([]int{})},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElements(tt.args.head, tt.args.val); !utils.IsLinkedListEqual(got, tt.want) {
				t.Errorf("removeElements() = %v, want %v", utils.ConvertLinkedListToSlice(got), utils.ConvertLinkedListToSlice(tt.want))
			}
		})
	}
}
