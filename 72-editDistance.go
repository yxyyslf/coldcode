//Given two strings word1 and word2, return the minimum number of operations
//required to convert word1 to word2.
//
// You have the following three operations permitted on a word:
//
//
// Insert a character
// Delete a character
// Replace a character
//
//
//
// Example 1:
//
//
//Input: word1 = "horse", word2 = "ros"
//Output: 3
//Explanation:
//horse -> rorse (replace 'h' with 'r')
//rorse -> rose (remove 'r')
//rose -> ros (remove 'e')
//
//
// Example 2:
//
//
//Input: word1 = "intention", word2 = "execution"
//Output: 5
//Explanation:
//intention -> inention (remove 't')
//inention -> enention (replace 'i' with 'e')
//enention -> exention (replace 'n' with 'x')
//exention -> exection (replace 'n' with 'c')
//exection -> execution (insert 'u')
//
//
//
// Constraints:
//
//
// 0 <= word1.length, word2.length <= 500
// word1 and word2 consist of lowercase English letters.
//
//
// Related Topics String Dynamic Programming 👍 9487 👎 107

package main

//leetcode submit region begin(Prohibit modification and deletion)
func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	distanceDp := make([][]int, m+1)
	for i := 0; i < len(distanceDp); i++ {
		distanceDp[i] = make([]int, n+1)
	}
	for i := 1; i <= n; i++ {
		distanceDp[0][i] = i
	}
	for i := 1; i <= m; i++ {
		distanceDp[i][0] = i
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				distanceDp[i][j] = min(min(distanceDp[i-1][j]+1, distanceDp[i][j-1]+1), distanceDp[i-1][j-1])
			} else {
				distanceDp[i][j] = 1 + min(min(distanceDp[i-1][j], distanceDp[i][j-1]), distanceDp[i-1][j-1])
			}
		}
	}
	return distanceDp[m][n]
}

//leetcode submit region end(Prohibit modification and deletion)
