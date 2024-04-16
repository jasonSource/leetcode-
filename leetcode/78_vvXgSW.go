package leetcode

func mergeKLists( lists []*ListNode ) *ListNode {
    // write code here
	if len(lists) == 0{
		return nil
	}
	if len(lists) == 0{
		return lists[0]
	}
	res := &ListNode{}
	for i :=0;i<len(lists);i++{
		res = mergeTwoLists1(res,lists[i])
	}
	return res
}


 /**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 
 * @param pHead ListNode类 
 * @param k int整型 
 * @return ListNode类
*/
func FindKthToTail( pHead *ListNode ,  k int ) *ListNode {
    // write code here
	first := pHead
	for i:=0;i<k;i++{
		if first == nil{
			return nil
		}
		first = first.Next
	}
	curr := pHead
	for first != nil{
		first = first.Next
		curr = curr.Next
	}
	return curr
}