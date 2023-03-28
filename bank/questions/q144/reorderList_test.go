package main

import (
	"algorithm-golang/types"
	"algorithm-golang/utils"
	"testing"
)

func Test_reorderList_withArray(t *testing.T) {
	tests := []struct {
		name string
		head *types.ListNode
		want *types.ListNode
	}{
		{"01", utils.GenerateLinkedList([]int{1, 2, 3, 4, 5}), utils.GenerateLinkedList([]int{1, 5, 2, 4, 3})},
		{"02", utils.GenerateLinkedList([]int{1, 2, 3, 4}), utils.GenerateLinkedList([]int{1, 4, 2, 3})},
		{"03", utils.GenerateLinkedList([]int{1, 2}), utils.GenerateLinkedList([]int{1, 2})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reorderListWithArray(tt.head)
			if !utils.IsLinkedListEqual(tt.head, tt.want) {
				t.Errorf("got %v, want %v\n", utils.ConvertLinkedListToSlice(tt.head), utils.ConvertLinkedListToSlice(tt.want))
			}
		})
	}
}

func Test_recordList(t *testing.T) {
	tests := []struct {
		name string
		head *types.ListNode
		want *types.ListNode
	}{
		{"01", utils.GenerateLinkedList([]int{1, 2, 3, 4, 5}), utils.GenerateLinkedList([]int{1, 5, 2, 4, 3})},
		{"02", utils.GenerateLinkedList([]int{1, 2, 3, 4}), utils.GenerateLinkedList([]int{1, 4, 2, 3})},
		{"03", utils.GenerateLinkedList([]int{1, 2}), utils.GenerateLinkedList([]int{1, 2})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reorderList(tt.head)
			if !utils.IsLinkedListEqual(tt.head, tt.want) {
				t.Errorf("got %v, want %v\n", utils.ConvertLinkedListToSlice(tt.head), utils.ConvertLinkedListToSlice(tt.want))
			}
		})
	}
}

func Test_merge(t *testing.T) {
	type args struct {
		h1 *types.ListNode
		h2 *types.ListNode
	}
	tests := []struct {
		name string
		args args
		want *types.ListNode
	}{
		{"01", args{h1: utils.GenerateLinkedList([]int{1, 2, 3}),
			h2: utils.GenerateLinkedList([]int{4, 5, 6}),
		}, utils.GenerateLinkedList([]int{1, 4, 2, 5, 3, 6})},
		{"02", args{
			h1: utils.GenerateLinkedList([]int{3, 4, 4}),
			h2: utils.GenerateLinkedList([]int{5, 6}),
		}, utils.GenerateLinkedList([]int{3, 5, 4, 6, 4})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.h1, tt.args.h2)
			if !utils.IsLinkedListEqual(tt.args.h1, tt.want) {
				t.Errorf("merge() = %v, want %v", utils.ConvertLinkedListToSlice(tt.args.h1), utils.ConvertLinkedListToSlice(tt.want))
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	tests := []struct {
		name string
		head *types.ListNode
		want *types.ListNode
	}{
		{"01", utils.GenerateLinkedList([]int{1, 2, 3, 4, 5}), utils.GenerateLinkedList([]int{5, 4, 3, 2, 1})},
		{"02", utils.GenerateLinkedList([]int{1, 2, 3}), utils.GenerateLinkedList([]int{3, 2, 1})},
		{"03", utils.GenerateLinkedList([]int{8, 9, 5, 4, 7, 1, 2}), utils.GenerateLinkedList([]int{2, 1, 7, 4, 5, 9, 8})},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.head); !utils.IsLinkedListEqual(got, tt.want) {
				t.Errorf("reverse() = %v, want %v", utils.ConvertLinkedListToSlice(got), utils.ConvertLinkedListToSlice(tt.want))
			}

		})
	}
}
