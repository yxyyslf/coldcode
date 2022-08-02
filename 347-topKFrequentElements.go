//Given an integer array nums and an integer k, return the k most frequent
//elements. You may return the answer in any order.
//
//
// Example 1:
// Input: nums = [1,1,1,2,2,3], k = 2
//Output: [1,2]
//
// Example 2:
// Input: nums = [1], k = 1
//Output: [1]
//
//
// Constraints:
//
//
// 1 <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
// k is in the range [1, the number of unique elements in the array].
// It is guaranteed that the answer is unique.
//
//
//
// Follow up: Your algorithm's time complexity must be better than O(n log n),
//where n is the array's size.
//
// Related Topics Array Hash Table Divide and Conquer Sorting Heap (Priority
//Queue) Bucket Sort Counting Quickselect 👍 10296 👎 401

package main

import "container/heap"

//leetcode submit region begin(Prohibit modification and deletion)
type NumFrequent struct {
	Number int
	Times  int
}

type ItemHeap []NumFrequent

func (h ItemHeap) Len() int { return len(h) }

func (h ItemHeap) Less(i, j int) bool {
	if h[i].Times != h[j].Times {
		return h[i].Times < h[j].Times
	} else {
		return h[i].Number > h[j].Number
	}
}

func (h ItemHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *ItemHeap) Push(val interface{}) {
	*h = append(*h, val.(NumFrequent))
}

func (h *ItemHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	numMap := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		_, ok := numMap[nums[i]]
		if ok {
			numMap[nums[i]]++
		} else {
			numMap[nums[i]] = 1
		}
	}
	h := &ItemHeap{}
	for w, v := range numMap {
		heap.Push(h, NumFrequent{
			Number: w,
			Times:  v,
		})

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := make([]int, h.Len())
	for i := h.Len() - 1; i >= 0; i-- {
		item := heap.Pop(h).(NumFrequent)
		res[i] = item.Number
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
