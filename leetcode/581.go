package leetcode

import "math"

/*
Input: [2, 6, 4, 8, 10, 9, 15]
Output: 5
Explanation: You need to sort [6, 4, 8, 10, 9] in ascending order to make the whole array sorted in ascending order.
*/

func findUnsortedSubarray(nums []int) int {
	start, end := -1, -2
	min, max := nums[len(nums)-1], nums[0]
	n := len(nums)
	for i := 0; i < len(nums); i++ {
		min = int(math.Min(float64(nums[n-1-i]), float64(min)))
		max = int(math.Max(float64(nums[i]), float64(max)))
		if nums[i] < max {
			end = i
		}
		if nums[n-1-i] > min {
			start = n-1-i
		}
	}
	return end-start+1
}
