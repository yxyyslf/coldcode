//Given the head of a linked list, return the list after sorting it in
//ascending order.
//
//
// Example 1:
//
//
//Input: head = [4,2,1,3]
//Output: [1,2,3,4]
//
//
// Example 2:
//
//
//Input: head = [-1,5,3,4,0]
//Output: [-1,0,3,4,5]
//
//
// Example 3:
//
//
//Input: head = []
//Output: []
//
//
//
// Constraints:
//
//
// The number of nodes in the list is in the range [0, 5 * 10⁴].
// -10⁵ <= Node.val <= 10⁵
//
//
//
// Follow up: Can you sort the linked list in O(n logn) time and O(1) memory (i.
//e. constant space)?
//
// Related Topics Linked List Two Pointers Divide and Conquer Sorting Merge
//Sort 👍 7841 👎 246

package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	slow, fast := dummy, dummy
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	newHead := slow.Next
	slow.Next = nil

	headA := sortList(head)
	headB := sortList(newHead)

	return sortMerge(headA, headB)
}

func sortMerge(headA, headB *ListNode) *ListNode {
	if headA == nil && headB == nil {
		return nil
	}
	curA := headA
	curB := headB
	newHead := &ListNode{}
	cur := newHead
	for curA != nil && curB != nil {
		if curA.Val > curB.Val {
			cur.Next = curB
			curB = curB.Next
		} else {
			cur.Next = curA
			curA = curA.Next
		}
		cur = cur.Next
	}
	if curA != nil {
		cur.Next = curA
	} else {
		cur.Next = curB
	}
	return newHead.Next
}

//leetcode submit region end(Prohibit modification and deletion)
