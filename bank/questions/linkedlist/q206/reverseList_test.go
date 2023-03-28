package main

import (
	"algorithm-golang/types"
	"algorithm-golang/utils"
	"testing"
)

func Test_reverseList(t *testing.T) {
	tests := []struct {
		name string
		head *types.ListNode
		want *types.ListNode
	}{
		{"01", utils.GenerateLinkedList([]int{1, 2, 3, 4, 5}), utils.GenerateLinkedList([]int{5, 4, 3, 2, 1})},
		{"02", utils.GenerateLinkedList([]int{1, 2}), utils.GenerateLinkedList([]int{2, 1})},
		{"03", utils.GenerateLinkedList([]int{}), utils.GenerateLinkedList([]int{})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.head); !utils.IsLinkedListEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
