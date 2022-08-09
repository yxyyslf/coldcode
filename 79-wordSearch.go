//Given an m x n grid of characters board and a string word, return true if
//word exists in the grid.
//
// The word can be constructed from letters of sequentially adjacent cells,
//where adjacent cells are horizontally or vertically neighboring. The same letter
//cell may not be used more than once.
//
//
// Example 1:
//
//
//Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
// "ABCCED"
//Output: true
//
//
// Example 2:
//
//
//Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
// "SEE"
//Output: true
//
//
// Example 3:
//
//
//Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
// "ABCB"
//Output: false
//
//
//
// Constraints:
//
//
// m == board.length
// n = board[i].length
// 1 <= m, n <= 6
// 1 <= word.length <= 15
// board and word consists of only lowercase and uppercase English letters.
//
//
//
// Follow up: Could you use search pruning to make your solution faster with a
//larger board?
//
// Related Topics Array Backtracking Matrix 👍 10383 👎 390

package main

//leetcode submit region begin(Prohibit modification and deletion)
func exist(board [][]byte, word string) bool {
	if board == nil || len(board) == 0 {
		return false
	}
	m := len(board)
	n := len(board[0])
	ch := word[0]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == ch && existTraverse(i, j, m, n, 0, &board, word) {
				return true
			}
		}
	}
	return false
}

func existTraverse(i, j, m, n, index int, board *[][]byte, word string) bool {
	if i < 0 || i >= m || j < 0 || j >= n || index == len(word) || (*board)[i][j] != word[index] {
		return false
	}
	if (*board)[i][j] == word[index] && index == len(word)-1 {
		return true
	}
	temp := (*board)[i][j]
	(*board)[i][j] = '0'
	flag := existTraverse(i+1, j, m, n, index+1, board, word) || existTraverse(i, j+1, m, n, index+1, board, word) ||
		existTraverse(i-1, j, m, n, index+1, board, word) || existTraverse(i, j-1, m, n, index+1, board, word)
	(*board)[i][j] = temp
	return flag
}

//leetcode submit region end(Prohibit modification and deletion)
