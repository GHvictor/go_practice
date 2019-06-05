package leetcode

/**
Input:
[[1,1,0],
 [1,1,0],
 [0,0,1]]
Output: 2
Explanation:The 0th and 1st students are direct friends, so they are in a friend circle.
The 2nd student himself is in a friend circle. So return 2.
 */
var tm, tn int
var tdirection = [4][2]int {
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func findCircleNum(M [][]int) int {
	if M == nil || len(M) == 0 {
		return 0
	}
	tm = len(M)
	tn = len(M[0])
	count := 0
	for i := 0; i < tm; i++ {
		for j := 0; j < tn; j++ {
			if M[i][j] == 1 {
				tdfs(M, i, j)
				count++
			}
		}
	}
	return count
}

func tdfs(M [][]int, i int, j int) {
	if i < 0 || i >= m || j < 0 || j >= n || M[i][j] == 0 {
		return
	}
	M[i][j] = 0
	for _, v := range tdirection {
		tdfs(M, i+v[0], j+v[1])
	}
	return
}
