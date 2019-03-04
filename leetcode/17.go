package leetcode

/**
Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
 */

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	lm := make(map[string][]string)
	lm["2"] = []string{"a", "b", "c"}
	lm["3"] = []string{"d", "e", "f"}
	lm["4"] = []string{"g", "h", "i"}
	lm["5"] = []string{"j", "k", "l"}
	lm["6"] = []string{"m", "n", "o"}
	lm["7"] = []string{"p", "q", "r", "s"}
	lm["8"] = []string{"t", "u", "v"}
	lm["9"] = []string{"w", "x", "y", "z"}
	ret := [][]string{}
	for i, v := range digits {
		if len(ret) <= 0 {
			ret = append(ret, lm[string(v)])
		} else {
			ret = append(ret, []string{})
			for _, s := range ret[i-1] {
				for _, ll := range lm[string(v)] {
					tmp := s + ll
					ret[i] = append(ret[i], tmp)
				}
			}
		}

	}
	return ret[len(digits)-1]
}
