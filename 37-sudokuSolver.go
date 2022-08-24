//Write a program to solve a Sudoku puzzle by filling the empty cells.
//
// A sudoku solution must satisfy all of the following rules:
//
//
// Each of the digits 1-9 must occur exactly once in each row.
// Each of the digits 1-9 must occur exactly once in each column.
// Each of the digits 1-9 must occur exactly once in each of the 9 3x3 sub-
//boxes of the grid.
//
//
// The '.' character indicates empty cells.
//
//
// Example 1:
//
//
//Input: board = [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5
//",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".
//",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".
//","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5
//"],[".",".",".",".","8",".",".","7","9"]]
//Output: [["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4
//","8"],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3
//"],["4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],[
//"9","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3",
//"4","5","2","8","6","1","7","9"]]
//Explanation: The input board is shown above and the only valid solution is
//shown below:
//
//
//
//
//
// Constraints:
//
//
// board.length == 9
// board[i].length == 9
// board[i][j] is a digit or '.'.
// It is guaranteed that the input board has only one solution.
//
//
// Related Topics Array Backtracking Matrix 👍 6257 👎 171

package main

//leetcode submit region begin(Prohibit modification and deletion)
func solveSudoku(board [][]byte) {
	rowNum := [10][10]bool{}
	colNum := [10][10]bool{}
	boxNum := [3][3][10]bool{}

	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			num := board[i][j] - '0'
			if board[i][j] != '.' {
				rowNum[i][num] = true
				colNum[j][num] = true
				boxNum[i/3][j/3][num] = true
			}
		}
	}
	solveSudokuTraverse(0, 0, board, rowNum, colNum, boxNum)
}

func solveSudokuTraverse(i, j int, board [][]byte, rowNum, colNum [10][10]bool, boxNum [3][3][10]bool) bool {
	if j == len(board[0]) {
		j = 0
		i++
		if i == len(board) {
			return true
		}
	}

	if board[i][j] == '.' {
		for c := '1'; c <= '9'; c++ {
			if rowNum[i][c-'0'] || colNum[j][c-'0'] || boxNum[i/3][j/3][c-'0'] {
				continue
			}
			rowNum[i][c-'0'] = true
			colNum[j][c-'0'] = true
			boxNum[i/3][j/3][c-'0'] = true
			board[i][j] = byte(c)
			if solveSudokuTraverse(i, j+1, board, rowNum, colNum, boxNum) {
				return true
			}
			board[i][j] = '.'
			rowNum[i][c-'0'] = false
			colNum[j][c-'0'] = false
			boxNum[i/3][j/3][c-'0'] = false
		}
	} else {
		return solveSudokuTraverse(i, j+1, board, rowNum, colNum, boxNum)
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
