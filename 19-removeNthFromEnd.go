package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
		if head == nil {
			return nil
		}

		slow := head
		fast := head
		for i:=0; i < n;i++ {
			fast = fast.Next
		}
		if fast == nil {
			return head.Next
		}


		for fast != nil {
			slow = slow.Next
			fast = fast.Next
		}
		slow.Next = slow.Next.Next
		return head
}
