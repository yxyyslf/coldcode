package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{
		Next: head,
	}
	slow := dummy
	fast := dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
		if fast == nil {
			return head
		}
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
