package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want bool
	}{
		{"01", 121, true},
		{"02", -9988, false},
		{"03", 10, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindromeWithoutStr(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want bool
	}{
		{"01", 121, true},
		{"02", -9988, false},
		{"03", 10, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeWithoutStr(tt.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
