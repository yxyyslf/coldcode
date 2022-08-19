//You are given an array of integers nums, there is a sliding window of size k
//which is moving from the very left of the array to the very right. You can only
//see the k numbers in the window. Each time the sliding window moves right by one
//position.
//
// Return the max sliding window.
//
//
// Example 1:
//
//
//Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
//Output: [3,3,5,5,6,7]
//Explanation:
//Window position                Max
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
//
//
// Example 2:
//
//
//Input: nums = [1], k = 1
//Output: [1]
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
// 1 <= k <= nums.length
//
//
// Related Topics Array Queue Sliding Window Heap (Priority Queue) Monotonic
//Queue 👍 11283 👎 373

package main

//leetcode submit region begin(Prohibit modification and deletion)
type MonotonousQueue []int

func (m MonotonousQueue) Empty() bool {
	return len(m) == 0
}

func (m MonotonousQueue) Len() int {
	return len(m)
}

func (m MonotonousQueue) Front() int {
	return m[0]
}

func (m MonotonousQueue) Back() int {
	return m[m.Len()-1]
}

func (m *MonotonousQueue) Push(val int) {
	for !m.Empty() && val > m.Back() {
		*m = (*m)[:m.Len()-1]
	}
	*m = append(*m, val)
}

func (m *MonotonousQueue) Pop(val int) {
	if !m.Empty() && m.Front() == val {
		*m = (*m)[1:]
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	result := make([]int, 0)
	var queue MonotonousQueue
	queue = make([]int, 0)
	for i := 0; i < k; i++ {
		queue.Push(nums[i])
	}
	result = append(result, queue.Front())
	for i := k; i < len(nums); i++ {
		queue.Pop(nums[i-k])
		queue.Push(nums[i])
		result = append(result, queue.Front())
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
