package main

//Given the root of a binary tree, return the length of the diameter of the
//tree.
//
// The diameter of a binary tree is the length of the longest path between any
//two nodes in a tree. This path may or may not pass through the root.
//
// The length of a path between two nodes is represented by the number of edges
//between them.
//
//
// Example 1:
//
//
//Input: root = [1,2,3,4,5]
//Output: 3
//Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].
//
//
// Example 2:
//
//
//Input: root = [1,2]
//Output: 1
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [1, 10⁴].
// -100 <= Node.val <= 100
//
//
// Related Topics Tree Depth-First Search Binary Tree 👍 8821 👎 548

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var result int

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result = 0
	traverse(root)
	return result
}

func traverse(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var left, right int
	if root.Left != nil {
		left = traverse(root.Left)
	} else {
		left = 0
	}
	if root.Right != nil {
		right = traverse(root.Right)
	} else {
		right = 0
	}
	result = max(result, left+right)
	return max(left, right) + 1
}