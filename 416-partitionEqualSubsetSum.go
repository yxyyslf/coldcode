//Given a non-empty array nums containing only positive integers, find if the
//array can be partitioned into two subsets such that the sum of elements in both
//subsets is equal.
//
//
// Example 1:
//
//
//Input: nums = [1,5,11,5]
//Output: true
//Explanation: The array can be partitioned as [1, 5, 5] and [11].
//
//
// Example 2:
//
//
//Input: nums = [1,2,3,5]
//Output: false
//Explanation: The array cannot be partitioned into equal sum subsets.
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 200
// 1 <= nums[i] <= 100
//
//
// Related Topics Array Dynamic Programming 👍 8355 👎 133

package main

//leetcode submit region begin(Prohibit modification and deletion)
func canPartition(nums []int) bool {
	partitionDp := make([][]bool, len(nums))
	sum := 0
	maxNum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		maxNum = max(maxNum, nums[i])
	}
	if sum%2 != 0 || maxNum > sum/2 {
		return false
	}
	target := sum / 2
	for i := 0; i < len(nums); i++ {
		partitionDp[i] = make([]bool, target+1)
		partitionDp[i][0] = true
	}
	partitionDp[0][nums[0]] = true
	for i := 1; i < len(nums); i++ {
		for j := 1; j <= target; j++ {
			if j > nums[i] {
				partitionDp[i][j] = partitionDp[i-1][j] || partitionDp[i-1][j-nums[i]]
			} else {
				partitionDp[i][j] = partitionDp[i-1][j]
			}
		}
	}
	return partitionDp[len(nums)-1][target]
}

//leetcode submit region end(Prohibit modification and deletion)
