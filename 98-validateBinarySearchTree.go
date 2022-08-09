//Given the root of a binary tree, determine if it is a valid binary search
//tree (BST).
//
// A valid BST is defined as follows:
//
//
// The left subtree of a node contains only nodes with keys less than the
//node's key.
// The right subtree of a node contains only nodes with keys greater than the
//node's key.
// Both the left and right subtrees must also be binary search trees.
//
//
//
// Example 1:
//
//
//Input: root = [2,1,3]
//Output: true
//
//
// Example 2:
//
//
//Input: root = [5,1,4,null,null,3,6]
//Output: false
//Explanation: The root node's value is 5 but its right child's value is 4.
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [1, 10⁴].
// -2³¹ <= Node.val <= 2³¹ - 1
//
//
// Related Topics Tree Depth-First Search Binary Search Tree Binary Tree 👍 1117
//6 👎 951

package main

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return isValidBSTraverse(root, math.MinInt, math.MaxInt)
}

func isValidBSTraverse(root *TreeNode, minVal, maxVal int) bool {
	if root == nil {
		return true
	}
	if root.Val >= maxVal || root.Val <= minVal {
		return false
	}
	return isValidBSTraverse(root.Left, minVal, root.Val) && isValidBSTraverse(root.Right, root.Val, maxVal)
}

//leetcode submit region end(Prohibit modification and deletion)
