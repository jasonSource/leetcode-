package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	curr := &ListNode{}
	durr := curr
	for list1 != nil || list2 != nil{
		temp := &ListNode{}
		if list1 == nil&& list2 != nil{
			temp.Val = list2.Val
			list2 = list2.Next 
		}else if list1 != nil&& list2 == nil{
			temp.Val = list1.Val
			list1 = list1.Next
		}else{
			if list1.Val < list2.Val{
				temp.Val = list1.Val
				list1 = list1.Next
			}else{
				temp.Val = list2.Val
				list2 = list2.Next 
			}
		}
		curr.Next = temp
		curr = curr.Next
	}
	return durr.Next
 }