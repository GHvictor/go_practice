package leetcode


/**
Input:
11110
11010
11000
00000

Output: 1

Input:
11000
11000
00100
00011

Output: 3
 */
var m, n int
var direction = [4][2]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}


func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	count := 0
	m = len(grid)
	n = len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != byte('0') {
				dfs(grid, i, j)
				count++
			}
		}
	}
	return count

}

func dfs(grid [][]byte, i int, j int) {
	if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == byte('0') {
		return
	}
	grid[i][j] = byte('0')
	for _, v := range direction {
		dfs(grid, i+v[0], j+v[1])
	}
	return
}
