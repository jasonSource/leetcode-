package main

import (
	"fmt"
	"testing"
)

//实现 strStr() 函数。

//给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回  -1 。

func TestStrStr(t *testing.T) {
	haystack := "hel"
	needle := "ll"
	fmt.Println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	length := len(needle)
	for i := 0; i <= len(haystack)-length; i++ {
		if haystack[i:i+length] == needle {
			return i
		}
	}
	return -1
}
