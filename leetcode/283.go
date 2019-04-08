package leetcode

/**
Input: [0,1,0,3,12]
Output: [1,3,12,0,0]
 */

func moveZeroes(nums []int)  {
	mark := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			mark = append(mark, i)
		}
	}

	for i, j := 0, 0; i < len(nums); i, j = i+1, j+1 {
		if j < len(mark) {
			nums[i] = nums[mark[j]]
		} else {
			nums[i] = 0
		}

	}

}
