package main

import (
	"reflect"
	"testing"
)

func Test_pairSums(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{nums: []int{5, 6, 5}, target: 11}, [][]int{{5, 6}}},
		//{"2", args{nums: []int{5, 6, 5, 6}, target: 11}, [][]int{{5, 6}, {5, 6}}},
		{"3", args{nums: []int{3, 3, 7}, target: 10}, [][]int{{3, 7}}},
		{"4", args{nums: []int{5, 5, 7, 8}, target: 2}, [][]int{}},
		{"5", args{nums: []int{-6, -3, 1, 9, 9, 2, 5, 6, -6, -2}, target: 3}, [][]int{{-6, 9}, {-6, 9}, {-3, 6}, {-2, 5}, {1, 2}}},
		{"6", args{nums: []int{-2, -1, 0, 3, 5, 6, 7, 9, 13, 14}, target: 12}, [][]int{{-2, 14}, {-1, 13}, {3, 9}, {5, 7}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pairSums(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pairSums() = %v, want %v", got, tt.want)
			}
		})
	}
}
