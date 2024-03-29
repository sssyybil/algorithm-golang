package main

import (
	"algorithm-golang/types"
	"algorithm-golang/utils"
	"reflect"
	"testing"
)

func Test_sortListWithInsert(t *testing.T) {
	type args struct {
		head *types.ListNode
	}
	tests := []struct {
		name string
		args args
		want *types.ListNode
	}{
		// TODO: Add test cases.
		{"1", args{utils.GenerateLinkedList([]int{4, 2, 1, 3})}, utils.GenerateLinkedList([]int{1, 2, 3, 4})},
		{"2", args{utils.GenerateLinkedList([]int{-1, 5, 3, 4, 0})}, utils.GenerateLinkedList([]int{-1, 0, 3, 4, 5})},
		{"3", args{utils.GenerateLinkedList([]int{1, 1})}, utils.GenerateLinkedList([]int{1, 1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortListWithInsert(tt.args.head); !utils.IsLinkedListEqual(got, tt.want) {
				t.Errorf("sortListWithInsert() = %v, want %v", utils.ConvertLinkedListToSlice(got), utils.ConvertLinkedListToSlice(tt.want))
			}
		})
	}
}

func Test_sortListWithMerge(t *testing.T) {
	type args struct {
		head *types.ListNode
	}
	tests := []struct {
		name string
		args args
		want *types.ListNode
	}{
		// TODO: Add test cases.
		{"1", args{utils.GenerateLinkedList([]int{4, 2, 1, 3})}, utils.GenerateLinkedList([]int{1, 2, 3, 4})},
		{"2", args{utils.GenerateLinkedList([]int{-1, 5, 3, 4, 0})}, utils.GenerateLinkedList([]int{-1, 0, 3, 4, 5})},
		{"3", args{utils.GenerateLinkedList([]int{1, 1})}, utils.GenerateLinkedList([]int{1, 1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortListWithMerge(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortListWithMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}
