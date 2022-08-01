package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
	a := []int{1, 1, 1, 1, 1}
	r := findTargetSumWays(a, 3)
	fmt.Println(r)
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
