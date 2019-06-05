package leetcode

/**
Input: "abc"
Output: 3
Explanation: Three palindromic strings: "a", "b", "c".
*/

func countSubstrings(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		for j := i+1; j <= len(s); j++ {
			if isSub(s[i:j]) {
				sum++
			}
		}
	}
	return sum

}

func isSub(s string) bool {
	if len(s) == 1 {
		return true
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true

}