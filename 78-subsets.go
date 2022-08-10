//Given an integer array nums of unique elements, return all possible subsets (
//the power set).
//
// The solution set must not contain duplicate subsets. Return the solution in
//any order.
//
//
// Example 1:
//
//
//Input: nums = [1,2,3]
//Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
//
//
// Example 2:
//
//
//Input: nums = [0]
//Output: [[],[0]]
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 10
// -10 <= nums[i] <= 10
// All the numbers of nums are unique.
//
//
// Related Topics Array Backtracking Bit Manipulation 👍 11380 👎 168

package main

//leetcode submit region begin(Prohibit modification and deletion)
func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	list := make([]int, 0)
	subsetsTraverse(0, nums, &list, &result)
	return result
}

func subsetsTraverse(i int, nums []int, list *[]int, result *[][]int) {
	if i == len(nums) {
		*result = append(*result, append([]int(nil), *list...))
		return
	}
	*list = append(*list, nums[i])
	subsetsTraverse(i+1, nums, list, result)
	*list = (*list)[:len(*list)-1]
	subsetsTraverse(i+1, nums, list, result)
}

//leetcode submit region end(Prohibit modification and deletion)
