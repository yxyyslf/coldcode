//Given an array nums with n objects colored red, white, or blue, sort them in-
//place so that objects of the same color are adjacent, with the colors in the
//order red, white, and blue.
//
// We will use the integers 0, 1, and 2 to represent the color red, white, and
//blue, respectively.
//
// You must solve this problem without using the library's sort function.
//
//
// Example 1:
//
//
//Input: nums = [2,0,2,1,1,0]
//Output: [0,0,1,1,2,2]
//
//
// Example 2:
//
//
//Input: nums = [2,0,1]
//Output: [0,1,2]
//
//
//
// Constraints:
//
//
// n == nums.length
// 1 <= n <= 300
// nums[i] is either 0, 1, or 2.
//
//
//
// Follow up: Could you come up with a one-pass algorithm using only constant
//extra space?
//
// Related Topics Array Two Pointers Sorting 👍 11426 👎 439

package main

//leetcode submit region begin(Prohibit modification and deletion)
func sortColors(nums []int) {
	leftIndex := 0
	rightIndex := len(nums)
	for _, num := range nums {
		if num == 0 {
			leftIndex++
		}
		if num == 2 {
			rightIndex--
		}
	}
	for i := range nums {
		if i < leftIndex {
			nums[i] = 0
		} else if i >= rightIndex {
			nums[i] = 2
		} else {
			nums[i] = 1
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
