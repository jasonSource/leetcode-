package main

import (
	"fmt"
	"strings"
	"testing"
)

//请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
//https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/

func TestReplaceSpace(t *testing.T) {
	s := "We are happy."
	result := replaceSpace(s)
	fmt.Println(result)
}
func replaceSpace(s string) string {
	var result string
	if len(s) == 0 {
		return result
	}
	for _, x := range strings.Split(s, "") {
		if x == " " {
			x = "%20"
		}
		result += x

	}
	return result
}
