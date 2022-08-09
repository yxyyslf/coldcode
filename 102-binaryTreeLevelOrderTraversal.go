//Given the root of a binary tree, return the level order traversal of its
//nodes' values. (i.e., from left to right, level by level).
//
//
// Example 1:
//
//
//Input: root = [3,9,20,null,null,15,7]
//Output: [[3],[9,20],[15,7]]
//
//
// Example 2:
//
//
//Input: root = [1]
//Output: [[1]]
//
//
// Example 3:
//
//
//Input: root = []
//Output: []
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [0, 2000].
// -1000 <= Node.val <= 1000
//
//
// Related Topics Tree Breadth-First Search Binary Tree 👍 10085 👎 195

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
func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0)

	nextQueue := make([]*TreeNode, 0)
	queue = append(queue, root)
	level := make([]int, 0)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			nextQueue = append(nextQueue, node.Left)
		}
		if node.Right != nil {
			nextQueue = append(nextQueue, node.Right)
		}
		level = append(level, node.Val)
		if len(queue) == 0 {
			result = append(result, level)
			queue = nextQueue
			nextQueue = make([]*TreeNode, 0)
			level = make([]int, 0)
		}
	}
	return result

}

//leetcode submit region end(Prohibit modification and deletion)
