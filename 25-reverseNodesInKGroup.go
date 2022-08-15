//Given the head of a linked list, reverse the nodes of the list k at a time,
//and return the modified list.
//
// k is a positive integer and is less than or equal to the length of the
//linked list. If the number of nodes is not a multiple of k then left-out nodes, in
//the end, should remain as it is.
//
// You may not alter the values in the list's nodes, only nodes themselves may
//be changed.
//
//
// Example 1:
//
//
//Input: head = [1,2,3,4,5], k = 2
//Output: [2,1,4,3,5]
//
//
// Example 2:
//
//
//Input: head = [1,2,3,4,5], k = 3
//Output: [3,2,1,4,5]
//
//
//
// Constraints:
//
//
// The number of nodes in the list is n.
// 1 <= k <= n <= 5000
// 0 <= Node.val <= 1000
//
//
//
// Follow-up: Can you solve the problem in O(1) extra memory space?
//
// Related Topics Linked List Recursion 👍 8543 👎 527

package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{
		Next: head,
	}
	pre := dummy

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummy.Next
			}
		}
		next := tail.Next
		head, tail = headTailReverse(head, tail)
		pre.Next = head
		tail.Next = next
		head = next
		pre = tail
	}
	return dummy.Next
}

func headTailReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	cur := head
	pre := tail.Next
	for pre != tail {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return tail, head
}

//leetcode submit region end(Prohibit modification and deletion)
