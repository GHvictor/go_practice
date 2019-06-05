package leetcode

import "sort"

/**
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
 */

func nextPermutation(nums []int)  {
	for i := len(nums)-1; i >= 0; i-- {
		for j := len(nums)-1; j > i; j-- {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
				sort.Ints(nums[i+1:])
				return
			}
		}
	}
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func backtace(rest [][]int, temp []int, mark []bool, nums []int) {
	if len(temp) == len(nums) {
		rest = append(rest, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if mark[i] {
			continue
		}
		mark[i] = true
		temp = append(temp, nums[i])
		backtace(rest, temp, mark, nums)
		mark[i] = false
		temp = temp[:len(temp)-1]
	}

}
