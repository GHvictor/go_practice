package leetcode

import "math"

/*
Input: [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
 */
func lengthOfLIS(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	dp := make([]int, len(nums))
	ret := -1

	for i := 0; i < len(nums); i++ {
		max := 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				max = int(math.Max(float64(max), float64(dp[j]+1)))
			}
		}
		dp[i] = max
		if ret < dp[i] {
			ret = dp[i]
		}
	}
	return ret
}
