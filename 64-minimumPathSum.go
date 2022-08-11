//Given a m x n grid filled with non-negative numbers, find a path from top
//left to bottom right, which minimizes the sum of all numbers along its path.
//
// Note: You can only move either down or right at any point in time.
//
//
// Example 1:
//
//
//Input: grid = [[1,3,1],[1,5,1],[4,2,1]]
//Output: 7
//Explanation: Because the path 1 → 3 → 1 → 1 → 1 minimizes the sum.
//
//
// Example 2:
//
//
//Input: grid = [[1,2,3],[4,5,6]]
//Output: 12
//
//
//
// Constraints:
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 200
// 0 <= grid[i][j] <= 100
//
//
// Related Topics Array Dynamic Programming Matrix 👍 8196 👎 109

package main

//leetcode submit region begin(Prohibit modification and deletion)
func minPathSum(grid [][]int) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	pathDp := make([][]int, m)
	sum := 0
	for i := 0; i < m; i++ {
		pathDp[i] = make([]int, n)
		sum += grid[i][0]
		pathDp[i][0] = sum
	}
	sum = 0
	for i := 0; i < n; i++ {
		sum += grid[0][i]
		pathDp[0][i] = sum
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			pathDp[i][j] = min(pathDp[i-1][j], pathDp[i][j-1]) + grid[i][j]
		}
	}

	return pathDp[m-1][n-1]
}

//leetcode submit region end(Prohibit modification and deletion)
