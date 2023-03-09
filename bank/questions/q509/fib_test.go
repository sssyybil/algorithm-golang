package main

import (
	"testing"
	"time"
)

func Test_fibWithRecur(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"recur-01", 4, 3},
		{"recur-02", 10, 55},
		{"recur-03", 40, 102334155},
		{"recur-04", 100, 3736710778780434371},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			startTime := time.Now()
			got := fibWithRecur(tt.n)
			t.Logf("%s spend : %vs\n", tt.name, time.Now().Sub(startTime).Seconds())
			if got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fibWithRecurBetter(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"recur-01", 4, 3},
		{"recur-02", 10, 55},
		{"recur-03", 100, 3736710778780434371},
		{"recur-04", 1000, 817770325994397771},
		{"recur-05", 100000, 2754320626097736315},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 初始化数组
			record = make([]int, tt.n+1)

			startTime := time.Now()
			got := fibWithRecurBetter(tt.n)
			t.Logf("%s spend : %vs\n", tt.name, time.Now().Sub(startTime).Seconds())

			if got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fib(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"recur-01", 4, 3},
		{"recur-02", 10, 55},
		{"recur-03", 15, 610},
		{"recur-04", 22, 17711},
		{"recur-04", 30, 832040},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			startTime := time.Now()
			got := fib(tt.n)
			t.Logf("%s spend : %vs\n", tt.name, time.Now().Sub(startTime).Seconds())
			if got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
