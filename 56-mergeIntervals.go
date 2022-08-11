//Given an array of intervals where intervals[i] = [starti, endi], merge all
//overlapping intervals, and return an array of the non-overlapping intervals that
//cover all the intervals in the input.
//
//
// Example 1:
//
//
//Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
//Output: [[1,6],[8,10],[15,18]]
//Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
//
//
// Example 2:
//
//
//Input: intervals = [[1,4],[4,5]]
//Output: [[1,5]]
//Explanation: Intervals [1,4] and [4,5] are considered overlapping.
//
//
//
// Constraints:
//
//
// 1 <= intervals.length <= 10⁴
// intervals[i].length == 2
// 0 <= starti <= endi <= 10⁴
//
//
// Related Topics Array Sorting 👍 15378 👎 555

package main

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func merge(intervals [][]int) [][]int {
	result := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})

	interval := intervals[0]
	result = append(result, append([]int{}, interval...))
	for i := 1; i < len(intervals); i++ {
		interval = result[len(result)-1]
		if interval[1] < intervals[i][0] {
			result = append(result, append([]int{}, intervals[i]...))
		} else {
			result = result[:len(result)-1]
			if interval[1] <= intervals[i][1] {
				interval[1] = intervals[i][1]
			}
			result = append(result, append([]int{}, interval...))
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
