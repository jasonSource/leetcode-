package leetcode

import (
	"fmt"
	"testing"
)

func PrintNodeList(head *ListNode){
	for head != nil{
		fmt.Print(head.Val,",")
		head = head.Next
	}
}

func TestReverseList(t *testing.T){
	head1 := &ListNode{Val: 1}
	head2 := &ListNode{Val: 2}
	head3 := &ListNode{Val: 3}
	head1.Next = head2
	head2.Next = head3
	PrintNodeList(head1)
	h1 := reverseList(head1)
	PrintNodeList(h1)
}

func TestDeleteDuplicates(t *testing.T){
	head1 := &ListNode{Val: 1}
	head2 := &ListNode{Val: 1}
	head3 := &ListNode{Val: 2}
	head1.Next = head2
	head2.Next = head3
	res := deleteDuplicates(head1)
	fmt.Println(res)

	
}