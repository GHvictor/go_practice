package leetcode

import (
	"container/list"
	"strconv"
	"strings"
	"unicode"
)

/*
s = "3[a]2[bc]", return "aaabcbc".
s = "3[a2[c]]", return "accaccacc".
s = "2[abc]3[cd]ef", return "abcabccdcdcdef".
 */

func revertString(str string) string {
	s := []byte(str)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}

func decodeString(s string) string {
	var ret string
	stack := list.New()
	for _, v := range s {
		if v != ']' {
			stack.PushBack(v)
		} else {
			tmp_str := ""
			num_str := ""
			count := 0
			for stack.Len() > 0 {
				e := stack.Back()
				t_v := e.Value.(rune)

				if t_v != '[' {
					if count == 0 {
						tmp_str += string(t_v)
						//fmt.Println("a: ", tmp_str)
					}
					if count == 1 && unicode.IsDigit(t_v) {
						num_str += string(t_v)
						//fmt.Println("b: ", num_str)
					} else if count == 1 && unicode.IsLetter(t_v) {
						//fmt.Println("c: ", string(t_v))
						break
					}
					stack.Remove(e)
				} else {
					if count == 1 {
						break
					} else {
						stack.Remove(e)
						count++
					}
				}
			}
			tmp_str = revertString(tmp_str)
			num_str = revertString(num_str)
			c, _ := strconv.Atoi(num_str)
			ttt := strings.Repeat(tmp_str, c)
			//fmt.Println(ttt)
			for _, sss := range ttt {
				stack.PushBack(sss)
			}
			/*
			for q := stack.Front(); q != nil; q = q.Next() {
				ret += string(q.Value.(rune))
			}
			fmt.Println(ret)
			*/
		}

	}
	for e := stack.Front(); e != nil; e = e.Next() {
		ret += string(e.Value.(rune))
	}
	return ret

}
