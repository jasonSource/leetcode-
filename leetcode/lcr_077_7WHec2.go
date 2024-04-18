package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	return sort1(head, nil)
}

func sort1(head1, tail *ListNode) *ListNode {
	if head1 == nil {
		return head1
	}
	if head1.Next == tail {
		head1.Next = nil
		return head1
	}
	slow, fast := head1, head1
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	mid := slow
	return mergeTwoLists1(sort1(head1, mid), sort1(mid, tail))
}

func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	curr := &ListNode{}
	durr := curr
	for list1 != nil || list2 != nil {
		temp := &ListNode{}
		if list1 == nil && list2 != nil {
			temp.Val = list2.Val
			list2 = list2.Next
		} else if list1 != nil && list2 == nil {
			temp.Val = list1.Val
			list1 = list1.Next
		} else {
			if list1.Val < list2.Val {
				temp.Val = list1.Val
				list1 = list1.Next
			} else {
				temp.Val = list2.Val
				list2 = list2.Next
			}
		}
		curr.Next = temp
		curr = curr.Next
	}
	return durr.Next
}
