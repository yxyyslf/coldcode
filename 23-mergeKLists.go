package main

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	} else if len(lists) == 2 {
		return mergeList(lists[0], lists[1])
	}
	l := 0
	r := len(lists) - 1
	mid := l + (r-l) >> 2
	left := mergeKLists(lists[l:mid])
	right := mergeKLists(lists[mid:])
	return mergeList(left, right)
}

func mergeList(headA *ListNode, headB *ListNode) *ListNode {
	if headA == nil && headB == nil {
		return nil
	}
	curA := headA
	curB := headB
	dummy := &ListNode{}
	newHead := dummy
	for curA != nil && curB != nil {
		node := &ListNode{}
		if curA.Val >= curB.Val {
			node.Val = curB.Val
			curB = curB.Next
		} else  {
			node.Val = curA.Val
			curA = curA.Next
		}
		dummy.Next = node
		dummy = dummy.Next
	}

	if curA != nil {
		dummy.Next = curA
	}
	if curB != nil {
		dummy.Next = curB
	}
	return newHead.Next
}
