package main

import (
	"algorithm-golang/types"
	"algorithm-golang/utils"
	"testing"
)

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *types.ListNode
	}
	tests := []struct {
		name string
		args args
		want *types.ListNode
	}{
		// TODO: Add test cases.
		{"1", args{head: utils.CreateLinkedList([]int{1, 2, 3, 4})}, utils.CreateLinkedList([]int{2, 1, 4, 3})},
		{"2", args{head: utils.CreateLinkedList([]int{4, 5, 6})}, utils.CreateLinkedList([]int{5, 4, 6})},
		{"3", args{head: utils.CreateLinkedList([]int{})}, utils.CreateLinkedList([]int{})},
		{"4", args{head: utils.CreateLinkedList([]int{1})}, utils.CreateLinkedList([]int{1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPairs(tt.args.head); !utils.IsLinkedListEqual(got, tt.want) {
				t.Errorf("swapPairs() = %v, want %v", utils.ConvertLinkedListToSlice(got), utils.ConvertLinkedListToSlice(tt.want))
			}
		})
	}
}
