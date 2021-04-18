package main

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

func TestReversePrint(t *testing.T) {
	var head ListNode
	var node1 ListNode
	var node2 ListNode
	head.Val = 1
	node1.Val = 3
	node2.Val = 2
	head.Next = &node1
	node1.Next = &node2
	arr := reversePrint(&head)
	fmt.Println(arr)
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	var count int
	for fast := head; fast != nil; fast = fast.Next {
		count++
	}
	arr := make([]int, count)
	count--
	for head != nil {
		arr[count] = head.Val
		head = head.Next
		count--
	}
	return arr
}
