//Given an array nums of distinct integers, return all the possible
//permutations. You can return the answer in any order.
//
//
// Example 1:
// Input: nums = [1,2,3]
//Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
// Example 2:
// Input: nums = [0,1]
//Output: [[0,1],[1,0]]
//
// Example 3:
// Input: nums = [1]
//Output: [[1]]
//
//
// Constraints:
//
//
// 1 <= nums.length <= 6
// -10 <= nums[i] <= 10
// All the integers of nums are unique.
//
//
// Related Topics Array Backtracking 👍 12227 👎 211

package main

//leetcode submit region begin(Prohibit modification and deletion)
func permute(nums []int) [][]int {
	result := make([][]int, 0)
	permuteTraverse(0, &nums, &result)
	return result
}

func permuteTraverse(index int, nums *[]int, result *[][]int) {
	if index == len(*nums) {
		*result = append(*result, append([]int{}, *nums...))
		return
	}

	for i := index; i < len(*nums); i++ {
		(*nums)[index], (*nums)[i] = (*nums)[i], (*nums)[index]
		permuteTraverse(index+1, nums, result)
		(*nums)[index], (*nums)[i] = (*nums)[i], (*nums)[index]
	}
}

//leetcode submit region end(Prohibit modification and deletion)
