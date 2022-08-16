//Given an integer n, return the number of structurally unique BST's (binary
//search trees) which has exactly n nodes of unique values from 1 to n.
//
//
// Example 1:
//
//
//Input: n = 3
//Output: 5
//
//
// Example 2:
//
//
//Input: n = 1
//Output: 1
//
//
//
// Constraints:
//
//
// 1 <= n <= 19
//
//
// Related Topics Math Dynamic Programming Tree Binary Search Tree Binary Tree ?
//? 7844 👎 313

package main

//leetcode submit region begin(Prohibit modification and deletion)
func numTrees(n int) int {
	G := make([]int, n+1)
	G[0] = 1
	G[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			F := G[j-1] * G[i-j]
			G[i] += F
		}
	}
	return G[n]
}

//leetcode submit region end(Prohibit modification and deletion)
