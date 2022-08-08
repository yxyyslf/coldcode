//Given an m x n 2D binary grid grid which represents a map of '1's (land) and
//'0's (water), return the number of islands.
//
// An island is surrounded by water and is formed by connecting adjacent lands
//horizontally or vertically. You may assume all four edges of the grid are all
//surrounded by water.
//
//
// Example 1:
//
//
//Input: grid = [
//  ["1","1","1","1","0"],
//  ["1","1","0","1","0"],
//  ["1","1","0","0","0"],
//  ["0","0","0","0","0"]
//]
//Output: 1
//
//
// Example 2:
//
//
//Input: grid = [
//  ["1","1","0","0","0"],
//  ["1","1","0","0","0"],
//  ["0","0","1","0","0"],
//  ["0","0","0","1","1"]
//]
//Output: 3
//
//
//
// Constraints:
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 300
// grid[i][j] is '0' or '1'.
//
//
// Related Topics Array Depth-First Search Breadth-First Search Union Find
//Matrix 👍 15407 👎 362

package main

//leetcode submit region begin(Prohibit modification and deletion)
func numIslands(grid [][]byte) int {
	result := 0
	if grid == nil || len(grid) == 0 {
		return result
	}
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				result++
				landTraverse(i, j, m, n, &grid)
			}
		}
	}
	return result

}

func landTraverse(i, j, m, n int, grid *[][]byte) {
	if i < 0 || i >= m || j < 0 || j >= n || (*grid)[i][j] != '1' {
		return
	}
	(*grid)[i][j] = '2'
	landTraverse(i+1, j, m, n, grid)
	landTraverse(i-1, j, m, n, grid)
	landTraverse(i, j+1, m, n, grid)
	landTraverse(i, j-1, m, n, grid)
}

//leetcode submit region end(Prohibit modification and deletion)
