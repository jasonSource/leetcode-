package leetcode


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	curr := head
	num := head.Val
	for curr.Next != nil{
		if curr.Next.Val == num{
			curr.Next = curr.Next.Next
		}else{
			curr = curr.Next
			num = curr.Val
		}
	}
	return head
 }