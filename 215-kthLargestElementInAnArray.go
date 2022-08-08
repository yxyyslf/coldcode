//Given an integer array nums and an integer k, return the kᵗʰ largest element
//in the array.
//
// Note that it is the kᵗʰ largest element in the sorted order, not the kᵗʰ
//distinct element.
//
// You must solve it in O(n) time complexity.
//
//
// Example 1:
// Input: nums = [3,2,1,5,6,4], k = 2
//Output: 5
//
// Example 2:
// Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
//Output: 4
//
//
// Constraints:
//
//
// 1 <= k <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
//
//
// Related Topics Array Divide and Conquer Sorting Heap (Priority Queue)
//Quickselect 👍 11273 👎 570

package main

import "container/heap"

//leetcode submit region begin(Prohibit modification and deletion)
type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	var h IntHeap
	for i := range nums {
		heap.Push(&h, nums[i])
	}

	heap.Init(&h)
	for i := 1; i < k; i++ {
		heap.Pop(&h)
	}
	x := heap.Pop(&h)
	return x.(int)
}

//leetcode submit region end(Prohibit modification and deletion)
