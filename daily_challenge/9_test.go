package main

import (
	"fmt"
	"testing"
)

//回文数字

func TestIsPalindrome(t *testing.T) {
	x := 121
	result := isPalindrome(x)
	fmt.Println(result)
}

func isPalindrome(x int) bool {
	if x < 0 || x%10 == 0 {
		return false
	}
	var y int
	for x > y {
		y = y*10 + x%10
		x = x / 10
	}
	if x == y || x == y/10 {
		return true
	}
	return false
}
