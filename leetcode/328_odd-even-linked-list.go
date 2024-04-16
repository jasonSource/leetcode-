package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func oddEvenList(head *ListNode) *ListNode {
	if head ==nil{
		return head
	}
	evenNode := head.Next
	odd := head
	even := evenNode
	for even != nil && even.Next != nil{
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenNode
	return head
 }