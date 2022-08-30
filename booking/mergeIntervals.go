package booking

import "sort"

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
//leetcode submit region begin(Prohibit modification and deletion)
func merge(intervals [][]int) [][]int {
	if intervals == nil || len(intervals) == 0 {
		return nil
	}
	result := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	result = append(result, intervals[0])
	for i := 1; i < len(intervals); i++ {
		pre := result[len(result)-1]
		cur := intervals[i]
		if pre[1] < cur[0] {
			result = append(result, cur)
		} else {
			if pre[1] >= cur[1] {
				continue
			} else {
				result = result[:len(result)-1]
				cur[0] = pre[0]
				result = append(result, cur)
			}
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

