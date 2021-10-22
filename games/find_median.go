package games

import (
	"sort"
)

func FindMedian(nums []int) int {
	sort.Ints(nums)
	var result int
	length := len(nums)
	middle := length / 2
	remainder := length % 2
	if remainder == 0 {
		result = (nums[middle-1] + nums[middle]) / 2
	} else {
		result = nums[middle]
	}
	return result
}
