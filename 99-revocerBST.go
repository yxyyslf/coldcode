//You are given the root of a binary search tree (BST), where the values of
//exactly two nodes of the tree were swapped by mistake. Recover the tree without
//changing its structure.
//
//
// Example 1:
//
//
//Input: root = [1,3,null,null,2]
//Output: [3,1,null,null,2]
//Explanation: 3 cannot be a left child of 1 because 3 > 1. Swapping 1 and 3
//makes the BST valid.
//
//
// Example 2:
//
//
//Input: root = [3,1,4,null,null,2]
//Output: [2,1,4,null,null,3]
//Explanation: 2 cannot be in the right subtree of 3 because 2 < 3. Swapping 2
//and 3 makes the BST valid.
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [2, 1000].
// -2³¹ <= Node.val <= 2³¹ - 1
//
//
//
//Follow up: A solution using
//O(n) space is pretty straight-forward. Could you devise a constant
//O(1) space solution?
//
// Related Topics Tree Depth-First Search Binary Search Tree Binary Tree 👍 6013
// 👎 194

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
var a *TreeNode
var b *TreeNode
var pre *TreeNode

func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}
	a = nil
	b = nil
	pre = nil
	recoverTreeTraverse(root)
	temp := a.Val
	a.Val = b.Val
	b.Val = temp
	return
}
func recoverTreeTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	recoverTreeTraverse(root.Left)
	if pre != nil && pre.Val > root.Val {
		b = root
		if a == nil {
			a = pre
		}
	}
	pre = root
	recoverTreeTraverse(root.Right)

}

//leetcode submit region end(Prohibit modification and deletion)
