//You are given an m x n grid where each cell can have one of three values:
//
//
// 0 representing an empty cell,
// 1 representing a fresh orange, or
// 2 representing a rotten orange.
//
//
// Every minute, any fresh orange that is 4-directionally adjacent to a rotten
//orange becomes rotten.
//
// Return the minimum number of minutes that must elapse until no cell has a
//fresh orange. If this is impossible, return -1.
//
//
// Example 1:
//
//
//Input: grid = [[2,1,1],[1,1,0],[0,1,1]]
//Output: 4
//
//
// Example 2:
//
//
//Input: grid = [[2,1,1],[0,1,1],[1,0,1]]
//Output: -1
//Explanation: The orange in the bottom left corner (row 2, column 0) is never
//rotten, because rotting only happens 4-directionally.
//
//
// Example 3:
//
//
//Input: grid = [[0,2]]
//Output: 0
//Explanation: Since there are already no fresh oranges at minute 0, the answer
//is just 0.
//
//
//
// Constraints:
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 10
// grid[i][j] is 0, 1, or 2.
//
//
// Related Topics Array Breadth-First Search Matrix 👍 7962 👎 297

package main

//leetcode submit region begin(Prohibit modification and deletion)
type Location struct {
	X int
	Y int
}

func orangesRotting(grid [][]int) int {
	queue := make([]Location, 0)
	directions := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	m := len(grid)
	n := len(grid[0])
	fresh := 0
	// init
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				loc := Location{
					X: i,
					Y: j,
				}
				queue = append(queue, loc)
			}
			if grid[i][j] == 1 {
				fresh++
			}
		}
	}

	if fresh == 0 {
		return 0
	}

	count := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			loc := queue[0]
			for _, direction := range directions {
				newX := loc.X + direction[0]
				newY := loc.Y + direction[1]
				if newX < 0 || newX >= m || newY < 0 || newY >= n {
					continue
				}

				if grid[newX][newY] == 1 {
					queue = append(queue, Location{X: newX, Y: newY})
					grid[newX][newY] = 2
					fresh--
				}
			}
			queue = queue[1:]
		}
		count++
	}

	if fresh == 0 {
		return count - 1
	}
	return -1

}

//leetcode submit region end(Prohibit modification and deletion)
