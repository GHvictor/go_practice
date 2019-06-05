package leetcode



func checkInclusion(s1 string, s2 string) bool {
	mark := make(map[string]int)

	for _, v := range(s1) {
		if _, ok := mark[string(v)]; ok {
			mark[string(v)]+=1
		} else {
			mark[string(v)] = 1
		}
	}
	for i := 0; i <= len(s2)-len(s1); i++ {
		mark2 := make(map[string]int)
		for _, v := range(s2[i:i+len(s1)]) {
			if _, ok := mark2[string(v)]; ok {
				mark2[string(v)]+=1
			} else {
				mark2[string(v)] = 1
			}
		}
		if checknums(mark2, mark) {
			return true
		}
	}
	return false
}

func checknums(m1 map[string]int, m2 map[string]int) bool {
	for k, v := range(m1) {
		if _, ok := m2[string(k)]; ok && v == m2[string(k)] {
			continue
		} else {
			return false
		}
	}
	return true
}


/*
func checkInclusion(s1 string, s2 string) bool {

	mark := []bool {}
	for i := 0; i < len(s1); i++ {
		mark = append(mark, false)
	}
	strlist := []string{}
	strPermutation("", &strlist, s1, mark)
	for i := 0; i < len(strlist); i++ {
		if strings.Contains(s2, strlist[i]) {
			return true
		}
	}
	return false

}

func strPermutation(s string, strlist *[]string, s1 string, mark []bool) {
	if len(s) == len(s1) {
		*strlist = append(*strlist, s)
	}
	for i, v := range s1 {
		if mark[i] {
			continue
		}
		mark[i] = true
		s = s+string(v)
		strPermutation(s, strlist, s1, mark)
		mark[i] = false
		s = s[:len(s)-1]
	}

}
*/
