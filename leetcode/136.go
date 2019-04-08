package leetcode

/**
Input: [2,2,1]
Output: 1
Input: [4,1,2,1,2]
Output: 4
 */

func singleNumber(nums []int) int {
	mark := make(map[int]int)
	num := 0
	for _, v := range nums {
		if _, ok := mark[v]; ok {
			mark[v]++
		} else {
			mark[v] = 1
		}
	}
	for k, v := range mark {
		if v == 1 {
			num = k
			break
		}
	}
	return num
}
