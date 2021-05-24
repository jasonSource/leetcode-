package main

import "testing"

func TestDeleteNode(t *testing.T) {
	node1 := &ListNode{Val: 4}
	node2 := &ListNode{Val: 5}
	node3 := &ListNode{Val: 1}
	node4 := &ListNode{Val: 9}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

}

func deleteNode(head *ListNode, val int) *ListNode {
	first := head
	target := head
	if head.Val == val {
		return head.Next
	}
	for first != nil {
		if first.Val == val {
			target.Next = first.Next
		}
		target = first
		first = first.Next

	}
	return head
}
