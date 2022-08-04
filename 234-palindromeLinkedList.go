//Given the head of a singly linked list, return true if it is a palindrome.
//
//
// Example 1:
//
//
//Input: head = [1,2,2,1]
//Output: true
//
//
// Example 2:
//
//
//Input: head = [1,2]
//Output: false
//
//
//
// Constraints:
//
//
// The number of nodes in the list is in the range [1, 10⁵].
// 0 <= Node.val <= 9
//
//
//
//Follow up: Could you do it in
//O(n) time and
//O(1) space?
//
// Related Topics Linked List Two Pointers Stack Recursion 👍 10151 👎 613

package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	dummy := &ListNode{Val: 0, Next: head}
	slow := dummy
	fast := dummy
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	newHead := slow.Next
	var pre *ListNode
	pre = nil
	for newHead != nil {
		next := newHead.Next
		newHead.Next = pre
		pre = newHead
		newHead = next
	}
	cur := head
	for pre != nil {
		if pre.Val != cur.Val {
			return false
		}
		pre = pre.Next
		cur = cur.Next
	}
	return true

}

//leetcode submit region end(Prohibit modification and deletion)
