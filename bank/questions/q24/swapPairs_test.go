package main

import (
	"algorithm-golang/types"
	"algorithm-golang/utils"
	"testing"
)

func Test_swapPairs(t *testing.T) {
	tests := []struct {
		name string
		head *types.ListNode
		want *types.ListNode
	}{
		{"01", utils.GenerateLinkedList([]int{1, 2, 3, 4}), utils.GenerateLinkedList([]int{2, 1, 4, 3})},
		{"02", utils.GenerateLinkedList([]int{1, 2, 3, 4, 5}), utils.GenerateLinkedList([]int{2, 1, 4, 3, 5})},
		{"03", utils.GenerateLinkedList([]int{1}), utils.GenerateLinkedList([]int{1})},
		{"04", utils.GenerateLinkedList([]int{}), utils.GenerateLinkedList([]int{})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPairs(tt.head); !utils.IsLinkedListEqual(got, tt.want) {
				t.Errorf("swapPairs() = %v, want %v", utils.ConvertLinkedListToSlice(got), utils.ConvertLinkedListToSlice(tt.want))
			}
		})
	}
}
