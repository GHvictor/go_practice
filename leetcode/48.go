package leetcode

/**
Given input matrix =
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

rotate the input matrix in-place such that it becomes:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]
 */

func rotate(matrix [][]int)  {
	n := len(matrix)
	if n < 1 {
		return
	}
	for i := 0; i < n-1; i++ {
		for j := i; j < n-i-1; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[n-j][i]
			matrix[n-j][i] = matrix[n-i][n-j]
			matrix[n-i][n-j] = matrix[j][n-i]
			matrix[j][n-i] = tmp
		}
	}
}