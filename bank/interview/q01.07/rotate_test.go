package main

import (
	"algorithm-golang/utils"
	"testing"
)

func Test_rotate(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		res    [][]int
	}{
		{"01", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{"02", [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
			[][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.matrix)
			if !utils.IsMatrixEqual(tt.matrix, tt.res) {
				t.Errorf("matrix=%v, want=%v\n", tt.matrix, tt.res)
			}
		})
	}
}
