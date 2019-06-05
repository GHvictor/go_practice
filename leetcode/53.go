package leetcode

import "math"

/**
Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
 */

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums)+1)
	dp[len(nums)-1] = nums[len(nums)-1]
	max_num := dp[len(nums)-1]
	for i := len(nums)-2; i >= 0; i-- {
		dp[i] = int(math.Max(float64(nums[i]), float64(nums[i]+dp[i+1])))
		if max_num < dp[i] {
			max_num = dp[i]
		}
	}
	return max_num
}
