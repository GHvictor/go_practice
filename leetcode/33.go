package leetcode


func search(nums []int, target int) int {

	i, j := 0, len(nums)-1

	for i <= j {
		mid := i+ (j-i)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] >= nums[i] {
			if nums[i] <= target && target <= nums[mid] {
				j = mid-1
			} else {
				i= mid+1
			}
		} else if nums[mid] <= nums[j] {
			if nums[mid] <= target && target <= nums[j] {
				i = mid+1
			} else {
				j =mid-1
			}

		}
	}
	return -1
}