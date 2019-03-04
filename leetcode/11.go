package leetcode

import "math"

/*
Input: [1,8,6,2,5,4,8,3,7]
Output: 49
 */

func maxArea(height []int) int {
	max_num := -1
	for i := 0; i < len(height)-1; i++ {
		for j := i+1; j < len(height); j++ {
			min_num := int(math.Min(float64(height[i]), float64(height[j])))
			if min_num*(j-i) > max_num {
				max_num = min_num*(j-i)
			}
		}
	}
	return max_num
}
