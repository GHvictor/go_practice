package leetcode

import "fmt"

type Interval struct {
	Start int
	End int
}

func (self Interval) compare(other Interval) bool {
	if self.Start > other.Start {
		return true
	}
	return false
}

func swap(arr []Interval, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func QuickSort(arr []Interval) {
	if len(arr) < 1 {
		return
	}
	mid, i:= arr[0], 0
	head, tail := 0, len(arr) - 1
	flag := true
	for head < tail {
		if flag {
			if mid.compare(arr[tail]) {
				swap(arr, i, tail)
				i = tail
				flag = false
			} else {
				tail--
			}
		}
		if !flag {
			if arr[head].compare(mid) {
				swap(arr, i, head)
				i = head
				flag = true
			} else {
				head++
			}
		}
	}
	QuickSort(arr[0:i])
	QuickSort(arr[i+1:len(arr)])
}

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
func merge(intervals []Interval) []Interval {
	// test
	intervals = []Interval{{1, 4}, {2,6}, {15,18}, {8,10}}
	intervals = []Interval{{1,4}, {4,5}}
	intervals = []Interval{{1,4}, {2,3}}
	QuickSort(intervals)
	new_arr := []Interval{}
	tmp_arr := intervals[0]
	for i := 1; i < len(intervals) ; i++ {
		fmt.Println(intervals[i])
		if tmp_arr.End >= intervals[i].Start {
			if tmp_arr.End < intervals[i].End {
				tmp_arr.End = intervals[i].End
			}
		} else {
			new_arr = append(new_arr, tmp_arr)
			tmp_arr = intervals[i]
		}
		if i == len(intervals) -1 {
			new_arr = append(new_arr, tmp_arr)
		}
	}
	fmt.Println(intervals)
	fmt.Println(new_arr)
	return new_arr
}
