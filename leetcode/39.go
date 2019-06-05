package leetcode

import (
	"sort"
)

/**
Input: candidates = [2,3,6,7], target = 7,
A solution set is:
[
  [7],
  [2,2,3]
]
Input: candidates = [2,3,5], target = 8,
A solution set is:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]
 */

 /*
func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	sort.Ints(candidates)
	for i:= 0; i < len(candidates); i++ {
		temp_list := []int{}
		for j := i; j < len(candidates); j++ {
			sum := 0
			for {
				if sum + candidates[j] < target {
					temp_list = append(temp_list, candidates[j])
					sum += candidates[j]
				} else if sum + candidates[j] == target {
					result = append(result, temp_list)
					temp_list = temp_list[:len(temp_list)-1]
					break
				} else if sum + candidates[j] > target {
					temp_list = temp_list[:len(temp_list)-1]
					break
				}

			}
		}
	}
	return result
}
 */

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	ret := [][]int{}
	findSum(&ret, []int{}, candidates, target)
	return ret
}

func findSum(ret *[][]int, temp []int, candidates []int, target int) {
	sum := 0
	for _, v := range temp {
		sum += v
	}
	if sum == target {
		*ret = append(*ret, temp)
		return
	} else if sum > target {
		return
	}
	for i:= 0; i < len(candidates); i++ {
		temp = append(temp, candidates[i])
		findSum(ret, temp, candidates, target)
		temp = temp[0:len(temp)-1]
	}
	return
}