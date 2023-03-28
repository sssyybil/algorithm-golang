package main

import (
	"algorithm-golang/types"
	"algorithm-golang/utils"
	"testing"
)

func Test_hasCycle(t *testing.T) {

	tests := []struct {
		name string
		head *types.ListNode
		want bool
	}{
		{"01", utils.GenerateCycleLinkedList([]int{1, 2, 3, 4}), true},
		{"02", utils.GenerateCycleLinkedList([]int{-21, 10, 17, 8, 4, 26, 5, 35, 33, -7, -16, 27, -12, 6, 29, -12, 5, 9, 20, 14, 14, 2, 13, -24, 21, 23, -21, 5}), true},
		{"03", utils.GenerateLinkedList([]int{1, 4, 6, 9}), false},
		{"03", utils.GenerateLinkedList([]int{1, 1, 4, 6, 7, 7, 8, 8}), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test(t *testing.T) {
	linkedList := utils.GenerateCycleLinkedList([]int{1, 2, 3, 4, 5})
	utils.PrintLinkedList(linkedList)
}
