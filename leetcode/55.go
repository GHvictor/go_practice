package leetcode


/**
Input: [2,3,1,1,4]
Output: true
Input: [3,2,1,0,4]
Output: false
 */
func canJump(nums []int) bool {
	dp := make([]bool, len(nums))
	dp[len(nums)-1] = true
	for i := len(nums)-2; i >= 0; i-- {
		if nums[i] >= len(nums) - i - 1 {
			dp[i] = true
		} else {
			for j := i; j < i + nums[i]; j++ {
				if dp[j] {
					dp[i] = dp[j]
					break
				}
			}
		}
	}
	return dp[0]
}
