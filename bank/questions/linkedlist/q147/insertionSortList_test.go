package main

import (
	"algorithm-golang/types"
	"algorithm-golang/utils"
	"testing"
)

func Test_insertionSortList(t *testing.T) {
	type args struct {
		head *types.ListNode
	}
	tests := []struct {
		name string
		args args
		want *types.ListNode
	}{
		// TODO: Add test cases.
		{"1", args{utils.CreateLinkedList([]int{4, 2, 1, 3})}, utils.CreateLinkedList([]int{1, 2, 3, 4})},
		{"2", args{utils.CreateLinkedList([]int{-1, 5, 3, 4, 0})}, utils.CreateLinkedList([]int{-1, 0, 3, 4, 5})},
		{"3", args{utils.CreateLinkedList([]int{1, 1})}, utils.CreateLinkedList([]int{1, 1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertionSortList(tt.args.head); !utils.IsLinkedListEqual(got, tt.want) {
				t.Errorf("insertionSortList() = %v, want %v", utils.ConvertLinkedListToSlice(got), utils.ConvertLinkedListToSlice(tt.want))
			}
		})
	}
}
