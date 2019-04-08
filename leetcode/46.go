package leetcode

/**
排列数，回溯法求解
Input: [1,2,3]
Output:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
 */

func permute(nums []int) [][]int {
	ret := make([][]int, 0)
	n := len(nums)
	mask := make([]bool, n)
	for i := 0; i < n; i++ {
		mask[i] = true
	}
	permutations(&ret, []int{}, nums, mask)
	return ret
}

func permutations(ret *[][]int, tmp_slice []int, nums []int, mask []bool) {
	if len(nums) == len(tmp_slice) {
		cpy := make([]int, len(tmp_slice))
		copy(cpy, tmp_slice)
		*ret = append(*ret, cpy)
		return
	}
	for i := 0; i < len(nums); i++ {
		if mask[i] {
			mask[i] = false
			tmp_slice = append(tmp_slice, nums[i])
			permutations(ret, tmp_slice, nums, mask)
			tmp_slice = tmp_slice[:len(tmp_slice)-1]
			mask[i] = true
		}
	}
	return

}
