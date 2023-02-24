package main

import (
	"reflect"
	"testing"
)

func Test_countBits_A(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		// TODO: Add test cases.
		{"01", 2, []int{0, 1, 1}},
		{"02", 5, []int{0, 1, 1, 2, 1, 2}},
		{"03", 8, []int{0, 1, 1, 2, 1, 2, 2, 3, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBitsA(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBitsA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countBits_B(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		// TODO: Add test cases.
		{"01", 2, []int{0, 1, 1}},
		{"02", 5, []int{0, 1, 1, 2, 1, 2}},
		{"03", 8, []int{0, 1, 1, 2, 1, 2, 2, 3, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBitsB(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBitsA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countBits_C(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		// TODO: Add test cases.
		{"01", 2, []int{0, 1, 1}},
		{"02", 5, []int{0, 1, 1, 2, 1, 2}},
		{"03", 8, []int{0, 1, 1, 2, 1, 2, 2, 3, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBitsC(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBitsA() = %v, want %v", got, tt.want)
			}
		})
	}
}
