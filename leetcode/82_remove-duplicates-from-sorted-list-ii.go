package leetcode


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil{
		return head
	}
	dummcy := &ListNode{Val: 0,Next: head}
	curr := dummcy
	for curr.Next != nil && curr.Next.Next != nil{
		if curr.Next.Val == curr.Next.Next.Val{
			x := curr.Next.Val
			for curr.Next != nil && curr.Next.Val == x{
				curr.Next = curr.Next.Next
			}
		}else{
			curr = curr.Next
		}
	}
	return dummcy.Next
 }