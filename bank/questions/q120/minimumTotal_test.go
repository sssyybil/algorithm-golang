package main

import "testing"

func Test_minimumTotal(t *testing.T) {
	tests := []struct {
		name     string
		triangle [][]int
		want     int
	}{
		{"01", [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}, 11},
		{"02", [][]int{{-10}}, -10},
		{"03", [][]int{{2}, {4, 5}, {1, 6, 9}, {2, 3, 4, 6}, {8, 8, 6, 6, 5}}, 16},
		{"04", [][]int{{-1}, {2, 3}, {1, -1, -3}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTotal(tt.triangle); got != tt.want {
				t.Errorf("minimumTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}
