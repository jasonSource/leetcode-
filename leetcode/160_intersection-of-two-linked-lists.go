package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1 := headA
	headMap := make(map[*ListNode]struct{})
	for p1 != nil {
		headMap[p1] = struct{}{}
		p1 = p1.Next
	}
	p2 := headB
	for p2 != nil {
		if _, ok := headMap[p2]; ok {
			return p2
		}
		p2 = p2.Next
	}
	return nil
}
