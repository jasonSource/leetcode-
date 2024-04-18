package leetcode

func search(nums []int, target int) int {
	// write code here
	if len(nums) == 0 {
		return -1
	}
	if target < nums[0] || target > nums[len(nums)-1] {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left
		if target == nums[mid] {
			return mid
		} else if target > mid {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return -1
}
