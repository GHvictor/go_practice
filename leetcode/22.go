package leetcode

/**
[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
 */

func generateParenthesis(n int) []string {
	if n <= 0 {
		return []string{}
	}
	str_slice := []string{}
	backtrace(n, n, "", &str_slice)
	return str_slice
}

func backtrace(l int, r int, ss string, str *[]string) {
	if l == 0 && r == 0 {
		*str = append(*str, ss)
		return
	}
	if l > 0 {
		ss = ss + "("
		backtrace(l-1, r, ss, str)
	}
	if r > 0 && l < r {
		ss = ss + ")"
		backtrace(l, r-1, ss, str)
	}
	return
}

