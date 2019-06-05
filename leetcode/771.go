package leetcode

import "strings"

/**
Input: J = "aA", S = "aAAbbbb"
Output: 3
 */
func numJewelsInStones(J string, S string) int {
	sum := 0
	for _, v := range(J) {
		sum += strings.Count(S, string(v))
	}
	return sum
}
