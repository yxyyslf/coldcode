//Given an array nums of size n, return the majority element.
//
// The majority element is the element that appears more than ⌊n / 2⌋ times.
//You may assume that the majority element always exists in the array.
//
//
// Example 1:
// Input: nums = [3,2,3]
//Output: 3
//
// Example 2:
// Input: nums = [2,2,1,1,1,2,2]
//Output: 2
//
//
// Constraints:
//
//
// n == nums.length
// 1 <= n <= 5 * 10⁴
// -10⁹ <= nums[i] <= 10⁹
//
//
//
//Follow-up: Could you solve the problem in linear time and in
//O(1) space?
//
// Related Topics Array Hash Table Divide and Conquer Sorting Counting 👍 11132
//👎 368

package main

//leetcode submit region begin(Prohibit modification and deletion)
func majorityElement(nums []int) int {
	num := -1
	if nums == nil || len(nums) == 0 {
		return num
	}
	num = nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if num == nums[i] {
			count++
		} else {
			count--
		}
		if count == 0 {
			num = nums[i]
			count = 1
		}
	}
	return num
}

//leetcode submit region end(Prohibit modification and deletion)
