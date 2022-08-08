//Given an m x n binary matrix filled with 0's and 1's, find the largest square
//containing only 1's and return its area.
//
//
// Example 1:
//
//
//Input: matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1",
//"1"],["1","0","0","1","0"]]
//Output: 4
//
//
// Example 2:
//
//
//Input: matrix = [["0","1"],["1","0"]]
//Output: 1
//
//
// Example 3:
//
//
//Input: matrix = [["0"]]
//Output: 0
//
//
//
// Constraints:
//
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 300
// matrix[i][j] is '0' or '1'.
//
//
// Related Topics Array Dynamic Programming Matrix 👍 7733 👎 165

package main

//leetcode submit region begin(Prohibit modification and deletion)
func maximalSquare(matrix [][]byte) int {
	maxVal := 0
	if matrix == nil || len(matrix) == 0 {
		return 0
	}
	matrixDp := make([][]int, len(matrix))
	for i := 0; i < len(matrixDp); i++ {
		matrixDp[i] = make([]int, len(matrix[0]))
		matrixDp[i][0] = int(matrix[i][0] - '0')
		maxVal = max(maxVal, matrixDp[i][0])
	}
	for i := 0; i < len(matrixDp[0]); i++ {
		matrixDp[0][i] = int(matrix[0][i] - '0')
		maxVal = max(maxVal, matrixDp[0][i])
	}

	for i := 1; i < len(matrixDp); i++ {
		for j := 1; j < len(matrixDp[0]); j++ {
			if matrix[i][j] == '1' {
				matrixDp[i][j] = min(min(matrixDp[i-1][j], matrixDp[i][j-1]), matrixDp[i-1][j-1]) + 1
				maxVal = max(maxVal, matrixDp[i][j])
			} else {
				matrixDp[i][j] = 0
			}
		}
	}

	return maxVal * maxVal
}

//leetcode submit region end(Prohibit modification and deletion)
