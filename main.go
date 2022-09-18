package main

import (
	"awesomeProject/booking"
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

type RandomNode struct {
	Val    int
	Next   *RandomNode
	Random *RandomNode
}

func main() {
	//var pWord, nWord string
	//var m, n, k int
	//
	//reader := bufio.NewReader(os.Stdin) // 标准输入输出
	//pWord, _ = reader.ReadString('\n')  // 回车结束
	//
	//reader = bufio.NewReader(os.Stdin) // 标准输入输出
	//nWord, _ = reader.ReadString('\n') // 回车结束
	//
	//fmt.Scan(&m)
	//hotels := make([]int, m)
	//for i := 0; i < m; i++ {
	//	var id int
	//	fmt.Scan(&id)
	//	hotels[i] = id
	//}
	//fmt.Scan(&n)
	//reviews := make([]string, n)
	//for i := 0; i < n; i++ {
	//	var review string
	//	reader = bufio.NewReader(os.Stdin)  // 标准输入输出
	//	review, _ = reader.ReadString('\n') // 回车结束
	//	reviews[i] = review
	//}
	//
	//fmt.Scan(&k)
	//pWord := "breakfast beach citycenter location metro view staff price"
	//nWord := "not"
	//hotels := []int32{1, 2, 1, 1, 2}
	//reviews := []string{"This hotel has a nice view of the citycenter. The location is perfect.",
	//	"The breakfast is ok. Regarding location, it is quite far from citycenter but price is cheap so it is worth.",
	//	"Loc is excellent, 5 minutes from citycenter. There is also a metro station very close to the hotel.",
	//	"Good price but I couldn't take my dog and there were other guests with dogs!",
	//	"Very friendly staff and good cost-benefit ratio. Its location is a bit far from citycenter.",
	//}
	//k := int32(2)
	//ans := booking.AwardTopKHotels(pWord, nWord, hotels, reviews, k)

	fmt.Println(booking.KnightMove(5, 5))
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
