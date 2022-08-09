package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type RandomNode struct {
	Val    int
	Next   *RandomNode
	Random *RandomNode
}

func main() {
	//items := make([]uint32, 2)
	//items[0] = 100
	//items[1] = 200
	//
	//fmt.Println(items)
	//
	//items = append(items, 10)
	//fmt.Println(items)
	head1 := &ListNode{
		Val: 4,
	}
	head2 := &ListNode{
		Val: 2,
	}
	head3 := &ListNode{
		Val: 1,
	}
	head4 := &ListNode{
		Val: 3,
	}
	head1.Next = head2
	head2.Next = head3
	head3.Next = head4

	r := sortList(head1)
	for r != nil {
		//fmt.Println(r.Val)
		r = r.Next
	}
}

func RemoveDuplicateInt(intSlice []uint32) []uint32 {
	allKeys := make(map[uint32]bool)
	list := make([]uint32, 0)
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
