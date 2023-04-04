package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"01", 2, true},
		{"02", 12321, true},
		{"03", 11221, false},
		{"04", 78906, false},
		{"05", 11, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.n); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPrimeNumber(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"01", 2, true},
		{"02", 5, true},
		{"03", 8, false},
		{"04", 78906, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPrimeNumber(tt.n); got != tt.want {
				t.Errorf("isPrimeNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_primePalindrome(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"01", 2, 2},
		{"02", 6, 7},
		{"03", 8, 11},
		{"04", 5657834, 7014107},
		{"05", 45887963, 100030001},
		{"06", 36433452, 100030001},
		{"07", 100000000, 100030001},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := primePalindrome(tt.n); got != tt.want {
				t.Errorf("primePalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
