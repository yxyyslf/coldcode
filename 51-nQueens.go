//The n-queens puzzle is the problem of placing n queens on an n x n chessboard
//such that no two queens attack each other.
//
// Given an integer n, return all distinct solutions to the n-queens puzzle.
//You may return the answer in any order.
//
// Each solution contains a distinct board configuration of the n-queens'
//placement, where 'Q' and '.' both indicate a queen and an empty space, respectively.
//
//
// Example 1:
//
//
//Input: n = 4
//Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
//Explanation: There exist two distinct solutions to the 4-queens puzzle as
//shown above
//
//
// Example 2:
//
//
//Input: n = 1
//Output: [["Q"]]
//
//
//
// Constraints:
//
//
// 1 <= n <= 9
//
//
// Related Topics Array Backtracking 👍 7813 👎 182

package main

//leetcode submit region begin(Prohibit modification and deletion)
func solveNQueens(n int) [][]string {
	result := make([][]string, 0)
	place := make([][]int, n)
	for i := 0; i < n; i++ {
		place[i] = make([]int, n)
	}

	solveNQueensTraverse(0, n, &place, &result)

	return result
}

func solveNQueensTraverse(row, n int, place *[][]int, result *[][]string) {
	if row == n {
		insertPlace(*place, result)
		return
	}

	for col := 0; col < n; col++ {
		if !isPlaceValid(row, col, n, *place) {
			continue
		}

		(*place)[row][col] = 1
		solveNQueensTraverse(row+1, n, place, result)
		(*place)[row][col] = 0
	}

}

func insertPlace(place [][]int, result *[][]string) {
	queenMap := map[int]string{1: "Q", 0: "."}
	n := len(place)
	placeStr := make([]string, 0)
	for i := 0; i < n; i++ {
		str := ""
		for j := 0; j < n; j++ {
			str += queenMap[place[i][j]]
		}
		placeStr = append(placeStr, str)
	}
	*result = append(*result, placeStr)
}

func isPlaceValid(row, col, n int, place [][]int) bool {
	for k := 0; k < n; k++ {
		if place[row][k] == 1 {
			return false
		}
		if place[k][col] == 1 {
			return false
		}
	}
	i, j := row-1, col-1
	for i >= 0 && j >= 0 {
		if place[i][j] == 1 {
			return false
		}
		i--
		j--
	}

	i, j = row-1, col+1
	for i >= 0 && j < n {
		if place[i][j] == 1 {
			return false
		}
		i--
		j++
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
