package main

import "testing"

func Test_climbStairs(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"01", 2, 2},
		{"02", 5, 8},
		{"03", 15, 987},
		{"04", 30, 1346269},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_climbStairsWithRecur(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"01", 2, 2},
		{"02", 5, 8},
		{"03", 15, 987},
		{"04", 30, 1346269},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairsWithRecur(tt.n); got != tt.want {
				t.Errorf("climbStairsWithRecur() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_climbStairsWithMemo(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"01", 2, 2},
		{"02", 5, 8},
		{"03", 15, 987},
		{"04", 30, 1346269},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairsWithMemo(tt.n); got != tt.want {
				t.Errorf("climbStairsWithRecur() = %v, want %v", got, tt.want)
			}
		})
	}
}
