//Given an integer array nums, find the contiguous subarray (containing at
//least one number) which has the largest sum and return its sum.
//
// A subarray is a contiguous part of an array.
//
//
// Example 1:
//
//
//Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
//Output: 6
//Explanation: [4,-1,2,1] has the largest sum = 6.
//
//
// Example 2:
//
//
//Input: nums = [1]
//Output: 1
//
//
// Example 3:
//
//
//Input: nums = [5,4,-1,7,8]
//Output: 23
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
//
//
//
// Follow up: If you have figured out the O(n) solution, try coding another
//solution using the divide and conquer approach, which is more subtle.
//
// Related Topics Array Divide and Conquer Dynamic Programming 👍 23751 👎 1131

package main

//leetcode submit region begin(Prohibit modification and deletion)
func maxSubArray(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	preVal := nums[0]
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		if preVal < 0 {
			preVal = nums[i]
		} else {
			preVal += nums[i]
		}
		result = max(result, preVal)
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
