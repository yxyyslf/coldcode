//Given an array of integers nums and an integer k, return the total number of
//subarrays whose sum equals to k.
//
// A subarray is a contiguous non-empty sequence of elements within an array.
//
//
// Example 1:
// Input: nums = [1,1,1], k = 2
//Output: 2
//
// Example 2:
// Input: nums = [1,2,3], k = 3
//Output: 2
//
//
// Constraints:
//
//
// 1 <= nums.length <= 2 * 10⁴
// -1000 <= nums[i] <= 1000
// -10⁷ <= k <= 10⁷
//
//
// Related Topics Array Hash Table Prefix Sum 👍 14474 👎 447

package main

//leetcode submit region begin(Prohibit modification and deletion)
func subarraySum(nums []int, k int) int {
	count := 0
	if nums == nil || len(nums) == 0 {
		return count
	}
	numMap := make(map[int]int, 0)
	preSum := 0
	numMap[0] = 1
	for _, num := range nums {
		preSum += num
		v, ok := numMap[preSum-k]
		if ok {
			count += v

		}
		numMap[preSum] += 1
	}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)
