package main

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	inter := [][]int{{1, 3}, {6, 9}}
	newInter := []int{11, 12}
	result := insert(inter, newInter)
	fmt.Println(result)
}

func insert(intervals [][]int, newInterval []int) [][]int {
	left, right := newInterval[0], newInterval[1]
	ans := [][]int{}
	merged := false
	for _, interval := range intervals {
		if interval[0] > right {
			if !merged {
				ans = append(ans, []int{left, right})
			} else if interval[1] < left {
				ans = append(ans, interval)
			} else {
				left = min(left, interval)

			}
		}
		if !merged {
			ans = append(ans, []int{left, right})
		}
	}

	for _, interval := range intervals {
		if interval[0] > right {
			// 在插入区间的右侧且无交集
			if !merged {
				ans = append(ans, []int{left, right})
				merged = true
			}
			ans = append(ans, interval)
		} else if interval[1] < left {
			// 在插入区间的左侧且无交集
			ans = append(ans, interval)
		} else {
			// 与插入区间有交集，计算它们的并集
			left = min(left, interval[0])
			right = max(right, interval[1])
		}
	}
	if !merged {
		ans = append(ans, []int{left, right})
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
