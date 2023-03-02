package main

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"01", "abcabcbb", 3},
		{"02", "bbbbb", 1},
		{"03", "pwwkew", 3},
		{"04", " ", 1},
		{"05", "au", 2},
		{"06", longStr1, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLongestSubstring_B(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"01", "abcabcbb", 3},
		{"02", "bbbbb", 1},
		{"03", "pwwkew", 3},
		{"04", " ", 1},
		{"05", "au", 2},
		{"06", longStr1, 95},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstringB(tt.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
