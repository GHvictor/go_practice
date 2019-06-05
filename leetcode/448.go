package leetcode

import "math"

/**
Input:
[4,3,2,7,8,2,3,1]

Output:
[5,6]
 */

func findDisappearedNumbers(nums []int) []int {
	ret := []int{}
	for i := 0; i < len(nums); i++ {
		index := int(math.Abs(float64(nums[i]))) - 1
		if nums[index] > 0 {
			nums[index] *= -1
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			ret = append(ret, i+1)
		}
	}
	return ret
}