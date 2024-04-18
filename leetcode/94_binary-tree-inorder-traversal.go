package leetcode

func inorderTraversal(root *TreeNode) []int {
	var nums []int
	nums = add1(root,nums)
	return nums
 }

 func add1(root *TreeNode,nums []int)[]int{
	if root == nil{
		return nums
	}
	nums = add(root.Left,nums)
	nums = append(nums, root.Val)
	nums = add(root.Right,nums)
	return nums
 }