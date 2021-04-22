package main

import (
	"fmt"
	"strings"
	"testing"
)

//回文

func TestIsPalindrome1(t *testing.T) {
	s := "0P"
	result := isPalindrome1(s)
	fmt.Println(result)
}

func isPalindrome1(s string) bool {
	var s1 string
	for i := 0; i < len(s); i++ {
		if isnum(s[i]) {
			s1 += strings.ToLower(string(s[i]))
		}
	}
	n := len(s1) / 2
	if n == 0 {
		return true
	}
	for j := 0; j < n; j++ {
		if s1[j] != s1[len(s1)-j-1] {
			return false
		}
	}
	return true
}
func isnum(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9')
}
