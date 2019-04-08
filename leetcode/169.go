package leetcode

/**
Input: [3,2,3]
Output: 3

Input: [2,2,1,1,1,2,2]
Output: 2
 */

func majorityElement(nums []int) int {
	mark := make(map[int]int)
	half := len(nums)/2
	ret := 0
	for i := 0; i < len(nums); i++ {
		if _, ok := mark[nums[i]]; ok {
			mark[nums[i]]++

		} else {
			mark[nums[i]] = 1
		}
		if mark[nums[i]] > half {
			ret = nums[i]
			break
		}
	}
	return ret
}
