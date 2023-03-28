package main

import "testing"

func Test_numSquares(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"01", 8, 2},
		{"02", 12, 3},
		{"03", 6, 3},
		{"04", 100, 1},
		{"05", 300, 3},
		{"06", 666, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSquares(tt.n); got != tt.want {
				t.Errorf("numSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numSquaresWithRecur(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"01", 8, 2},
		{"02", 12, 3},
		{"03", 6, 3},
		{"04", 100, 1},
		{"05", 300, 3},
		{"06", 666, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSquaresWithRecur(tt.n); got != tt.want {
				t.Errorf("numSquaresWithRecur() = %v, want %v", got, tt.want)
			}
		})
	}
}
