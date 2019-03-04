package leetcode

/**
Input: "()"
Output: true
Input: "{[]}"
Output: true
( { [
 */

func isValid(s string) bool {
	hm := make(map[string]string)
	hm["{"] = "}"
	hm["["] = "]"
	hm["("] = ")"
	queue := []string{}
	for _, v := range s {
		if len(queue) <= 0 {
			queue = append(queue, string(v))
		} else {
			if hm[queue[len(queue)-1]] == string(v) {
				queue = queue[:len(queue)-1]
			} else {
				queue = append(queue, string(v))
			}
		}
	}
	if len(queue) == 0 {
		return true
	} else {
		return false
	}
}
