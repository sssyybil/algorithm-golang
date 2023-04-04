package main

import "testing"

func Test_encryptEmail(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"01", "LeetCode@LeetCode.com", "l*****e@leetcode.com"},
		{"02", "AB@qq.com", "a*****b@qq.com"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encryptEmail(tt.s); got != tt.want {
				t.Errorf("encryptEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encryptPhone(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"01", "1(234)567-890", "***-***-7890"},
		{"02", "86-(10)12345678", "+**-***-***-5678"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encryptPhone(tt.s); got != tt.want {
				t.Errorf("encryptPhone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maskPII(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"01", "LeetCode@LeetCode.com", "l*****e@leetcode.com"},
		{"02", "AB@qq.com", "a*****b@qq.com"},
		{"03", "1(234)567-890", "***-***-7890"},
		{"04", "86-(10)12345678", "+**-***-***-5678"}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maskPII(tt.s); got != tt.want {
				t.Errorf("maskPII() = %v, want %v", got, tt.want)
			}
		})
	}
}
