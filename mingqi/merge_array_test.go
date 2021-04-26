package mingqi

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	merge([]int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3)

}
func merge(nums1 []int, m int, nums2 []int, n int) {
	temp := make([]int, m)
	copy(temp, nums1)
	i, j, t := 0, 0, 0
	for ; t < len(nums1); t++ {
		if i == m && j < n {
			nums1[t] = nums2[j]
			j++
			continue
		}
		if i < m && j == n {
			nums1[t] = temp[i]
			i++
			continue
		}
		if nums2[j] >= temp[i] {
			nums1[t] = temp[i]
			i++
		} else {
			nums1[t] = nums2[j]
			j++
		}
	}
	fmt.Println(nums1)
}
