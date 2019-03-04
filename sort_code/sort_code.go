package sort_code

func compare(a int, b int) bool {
	if a > b {
		return true
	}
	return false
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func QuickSort(arr []int) {
	if len(arr) < 1 {
		return
	}
	mid, i := arr[0], 0
	head, tail := 0, len(arr) - 1
	flag := true
	for head < tail {
		if flag {
			if compare(mid, arr[tail]) {
				swap(arr, i, tail)
				i = tail
				flag = false
			} else {
				tail--
			}
		}
		if !flag {
			if compare(arr[head], mid) {
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
