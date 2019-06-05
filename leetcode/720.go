package leetcode

import (
	"sort"
)

func longestWord(words []string) string {
	sort.Strings(words)
	longest := ""
	hash_map := make(map[string]int)

	for i := 0; i < len(words); i++ {
		if len(words[i]) == 1 {
			if len(longest) < len(words[i]) {
				longest = words[i]
			}
			hash_map[words[i]] = 1
			continue
		}
		if _, ok := hash_map[words[i][:len(words[i])-1]]; ok {
			if len(longest) < len(words[i]) {
				longest = words[i]
			}
			hash_map[words[i]] = 1
		}
	}
	return longest
}