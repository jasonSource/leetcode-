package mingqi

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	test := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	result := maxSubArray(test)
	fmt.Println(result)
}
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
