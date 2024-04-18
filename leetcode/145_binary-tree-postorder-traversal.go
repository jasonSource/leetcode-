package leetcode

func postorderTraversal(root *TreeNode) []int {
	var nums []int
	nums = add2(root, nums)
	return nums
}

func add2(root *TreeNode, nums []int) []int {
	if root == nil {
		return nums
	}
	nums = add(root.Left, nums)
	nums = add(root.Right, nums)
	nums = append(nums, root.Val)
	return nums
}
