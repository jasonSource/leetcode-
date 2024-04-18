package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	var nums []int
	nums = add(root, nums)
	return nums
}

func add(root *TreeNode, nums []int) []int {
	if root == nil {
		return nums
	}
	nums = add(root.Left, nums)
	nums = append(nums, root.Val)
	nums = add(root.Right, nums)
	return nums
}
