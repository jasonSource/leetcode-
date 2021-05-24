package main

import (
	"fmt"
	"testing"
)

//给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字
type ListNode struct {
	Val  int
	Next *ListNode
}

func TestAddTwoNumber(t *testing.T) {
	node1 := &ListNode{Val: 2, Next: nil}
	node2 := &ListNode{Val: 4, Next: nil}
	node3 := &ListNode{Val: 3, Next: nil}
	node1.Next = node2
	node2.Next = node3
	node4 := &ListNode{Val: 5, Next: nil}
	node5 := &ListNode{Val: 6, Next: nil}
	//node6 := &ListNode{Val: 4,Next: nil}
	node4.Next = node5
	//node5.Next = node6
	l := addTwoNumber(node1, node4)
	printLink(l)

}

//遍历链表
func printLink(l1 *ListNode) {
	for l1 != nil {
		fmt.Println(l1.Val)
		l1 = l1.Next
	}
}

func addTwoNumber(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	tail := head
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return head
}
