package main

import "testing"

func Test_integerBreak(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"01", 2, 1},
		{"02", 10, 36},
		{"03", 15, 243},
		{"04", 33, 177147},
		{"05", 58, 1549681956},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := integerBreak(tt.n); got != tt.want {
				t.Errorf("integerBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_integerBreakWithDp(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"01", 2, 1},
		{"02", 10, 36},
		{"03", 15, 243},
		{"04", 33, 177147},
		{"05", 58, 1549681956},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := integerBreakWithDp(tt.n); got != tt.want {
				t.Errorf("integerBreakWithDp() = %v, want %v", got, tt.want)
			}
		})
	}
}
