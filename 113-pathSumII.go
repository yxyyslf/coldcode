//Given the root of a binary tree and an integer targetSum, return all root-to-
//leaf paths where the sum of the node values in the path equals targetSum. Each
//path should be returned as a list of the node values, not node references.
//
// A root-to-leaf path is a path starting from the root and ending at any leaf
//node. A leaf is a node with no children.
//
//
// Example 1:
//
//
//Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
//Output: [[5,4,11,2],[5,8,4,5]]
//Explanation: There are two paths whose sum equals targetSum:
//5 + 4 + 11 + 2 = 22
//5 + 8 + 4 + 5 = 22
//
//
// Example 2:
//
//
//Input: root = [1,2,3], targetSum = 5
//Output: []
//
//
// Example 3:
//
//
//Input: root = [1,2], targetSum = 0
//Output: []
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [0, 5000].
// -1000 <= Node.val <= 1000
// -1000 <= targetSum <= 1000
//
//
// Related Topics Backtracking Tree Depth-First Search Binary Tree 👍 5355 👎 11
//6

package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)
	pathSumTraverse(root, targetSum, &path, &result)
	return result
}

func pathSumTraverse(root *TreeNode, targetSum int, path *[]int, result *[][]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		*path = append(*path, root.Val)
		temp := make([]int, len(*path))
		copy(temp, *path)
		*result = append(*result, temp)
		*path = (*path)[:len(*path)-1]
		return
	}
	*path = append(*path, root.Val)
	pathSumTraverse(root.Left, targetSum-root.Val, path, result)
	pathSumTraverse(root.Right, targetSum-root.Val, path, result)
	*path = (*path)[:len(*path)-1]
}

//leetcode submit region end(Prohibit modification and deletion)
