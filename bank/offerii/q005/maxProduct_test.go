package main

import "testing"

func Test_maxProduct_A(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  int
	}{
		{"01", []string{"abcw", "baz", "foo", "bar", "fxyz", "abcdef"}, 16},
		{"02", []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}, 4},
		{"03", []string{"a", "aa", "aaa", "aaaa"}, 0},
		{"04", []string{"eae", "ea", "aaf", "bda", "fcf", "dc", "ac", "ce", "cefde", "dabae"}, 15},
		{"05", LongStr, 976144},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProductA(tt.words); got != tt.want {
				t.Errorf("maxProductA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProduct_B(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  int
	}{
		{"01", []string{"abcw", "baz", "foo", "bar", "fxyz", "abcdef"}, 16},
		{"02", []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}, 4},
		{"03", []string{"a", "aa", "aaa", "aaaa"}, 0},
		{"04", []string{"eae", "ea", "aaf", "bda", "fcf", "dc", "ac", "ce", "cefde", "dabae"}, 15},
		{"05", LongStr, 976144},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProductB(tt.words); got != tt.want {
				t.Errorf("maxProductA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProduct_C(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  int
	}{
		{"01", []string{"abcw", "baz", "foo", "bar", "fxyz", "abcdef"}, 16},
		{"02", []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}, 4},
		{"03", []string{"a", "aa", "aaa", "aaaa"}, 0},
		{"04", []string{"eae", "ea", "aaf", "bda", "fcf", "dc", "ac", "ce", "cefde", "dabae"}, 15},
		{"05", LongStr, 976144},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProductC(tt.words); got != tt.want {
				t.Errorf("maxProductA() = %v, want %v", got, tt.want)
			}
		})
	}
}
