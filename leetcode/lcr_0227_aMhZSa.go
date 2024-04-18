package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	slow, fast := head, head

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	second := reverse1(slow)
	for second != nil {
		if head.Val != second.Val {
			return false
		}
		head = head.Next
		second = second.Next
	}
	return true
}

func reverse1(head *ListNode) *ListNode {
	var pre *ListNode
	curr := head
	for curr != nil {
		temp := curr.Next
		curr.Next = pre
		pre = curr
		curr = temp
	}
	return pre
}
