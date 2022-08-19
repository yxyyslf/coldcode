//Given an integer array nums, find a contiguous non-empty subarray within the
//array that has the largest product, and return the product.
//
// The test cases are generated so that the answer will fit in a 32-bit integer.
//
//
// A subarray is a contiguous subsequence of the array.
//
//
// Example 1:
//
//
//Input: nums = [2,3,-2,4]
//Output: 6
//Explanation: [2,3] has the largest product 6.
//
//
// Example 2:
//
//
//Input: nums = [-2,0,-1]
//Output: 0
//Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 2 * 10⁴
// -10 <= nums[i] <= 10
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit
//integer.
//
//
// Related Topics Array Dynamic Programming 👍 12956 👎 392

package main

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
func maxProduct(nums []int) int {
	maxVal := 1
	minVal := 1
	result := math.MinInt
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			minVal, maxVal = maxVal, minVal
		}
		maxVal = max(maxVal*nums[i], nums[i])
		minVal = min(minVal*nums[i], nums[i])

		result = max(result, maxVal)
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
