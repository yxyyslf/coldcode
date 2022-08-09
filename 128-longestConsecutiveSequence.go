//Given an unsorted array of integers nums, return the length of the longest
//consecutive elements sequence.
//
// You must write an algorithm that runs in O(n) time.
//
//
// Example 1:
//
//
//Input: nums = [100,4,200,1,3,2]
//Output: 4
//Explanation: The longest consecutive elements sequence is [1, 2, 3, 4].
//Therefore its length is 4.
//
//
// Example 2:
//
//
//Input: nums = [0,3,7,2,5,8,4,6,0,1]
//Output: 9
//
//
//
// Constraints:
//
//
// 0 <= nums.length <= 10⁵
// -10⁹ <= nums[i] <= 10⁹
//
//
// Related Topics Array Hash Table Union Find 👍 12437 👎 524

package main

//leetcode submit region begin(Prohibit modification and deletion)
func longestConsecutive(nums []int) int {
	numMap := make(map[int]bool, 0)
	for _, num := range nums {
		numMap[num] = true
	}
	result := 0
	for _, num := range nums {
		if numMap[num-1] {
			continue
		}
		curNum := num
		curLen := 1
		for numMap[curNum+1] {
			curNum++
			curLen++
		}
		if curLen > result {
			result = curLen
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
