package games

import "sort"

func RemoveElement(nums []int, val int) int {
	var result int
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums[i] = -1
		} else {
			result++
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	return result
}
