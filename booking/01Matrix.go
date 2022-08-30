package booking

//Given an m x n binary matrix mat, return the distance of the nearest 0 for
//each cell.
//
// The distance between two adjacent cells is 1.
//
//
// Example 1:
//
//
//Input: mat = [[0,0,0],[0,1,0],[0,0,0]]
//Output: [[0,0,0],[0,1,0],[0,0,0]]
//
//
// Example 2:
//
//
//Input: mat = [[0,0,0],[0,1,0],[1,1,1]]
//Output: [[0,0,0],[0,1,0],[1,2,1]]
//
//
//
// Constraints:
//
//
// m == mat.length
// n == mat[i].length
// 1 <= m, n <= 10⁴
// 1 <= m * n <= 10⁴
// mat[i][j] is either 0 or 1.
// There is at least one 0 in mat.
//
//
// Related Topics Array Dynamic Programming Breadth-First Search Matrix 👍 5466
//👎 270

//leetcode submit region begin(Prohibit modification and deletion)

var dirX = []int{1, -1, 0, 0}
var dirY = []int{0, 0, 1, -1}

func updateMatrix(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])

	queue := make([]Location, 0)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				queue = append(queue, Location{
					X: i,
					Y: j,
				})
			} else {
				mat[i][j] = -1
			}
		}
	}

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			for j := range dirX {
				newX := cur.X + dirX[j]
				newY := cur.Y + dirY[j]
				if newX >= 0 && newX < m && newY >= 0 && newY < n && mat[newX][newY] == -1 {
					mat[newX][newY] = mat[cur.X][cur.Y] + 1
					queue = append(queue, Location{
						X: newX,
						Y: newY,
					})
				}

			}
		}
	}

	return mat
}

//leetcode submit region end(Prohibit modification and deletion)
