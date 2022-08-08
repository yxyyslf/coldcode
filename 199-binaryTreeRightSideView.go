//Given the root of a binary tree, imagine yourself standing on the right side
//of it, return the values of the nodes you can see ordered from top to bottom.
//
//
// Example 1:
//
//
//Input: root = [1,2,3,null,5,null,4]
//Output: [1,3,4]
//
//
// Example 2:
//
//
//Input: root = [1,null,3]
//Output: [1,3]
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
// The number of nodes in the tree is in the range [0, 100].
// -100 <= Node.val <= 100
//
//
// Related Topics Tree Depth-First Search Breadth-First Search Binary Tree 👍 80
//89 👎 469

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
func rightSideView(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	nextQueue := make([]*TreeNode, 0)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Left != nil {
			nextQueue = append(nextQueue, node.Left)
		}
		if node.Right != nil {
			nextQueue = append(nextQueue, node.Right)
		}

		if len(queue) == 0 {
			result = append(result, node.Val)
		}

		if len(nextQueue) > 0 && len(queue) == 0 {
			queue = nextQueue
			nextQueue = make([]*TreeNode, 0)
		}
	}
	return result

}

//leetcode submit region end(Prohibit modification and deletion)
