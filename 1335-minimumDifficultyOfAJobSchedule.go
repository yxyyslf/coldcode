//You want to schedule a list of jobs in d days. Jobs are dependent (i.e To
//work on the iᵗʰ job, you have to finish all the jobs j where 0 <= j < i).
//
// You have to finish at least one task every day. The difficulty of a job
//schedule is the sum of difficulties of each day of the d days. The difficulty of a
//day is the maximum difficulty of a job done on that day.
//
// You are given an integer array jobDifficulty and an integer d. The
//difficulty of the iᵗʰ job is jobDifficulty[i].
//
// Return the minimum difficulty of a job schedule. If you cannot find a
//schedule for the jobs return -1.
//
//
// Example 1:
//
//
//Input: jobDifficulty = [6,5,4,3,2,1], d = 2
//Output: 7
//Explanation: First day you can finish the first 5 jobs, total difficulty = 6.
//Second day you can finish the last job, total difficulty = 1.
//The difficulty of the schedule = 6 + 1 = 7
//
//
// Example 2:
//
//
//Input: jobDifficulty = [9,9,9], d = 4
//Output: -1
//Explanation: If you finish a job per day you will still have a free day. you
//cannot find a schedule for the given jobs.
//
//
// Example 3:
//
//
//Input: jobDifficulty = [1,1,1], d = 3
//Output: 3
//Explanation: The schedule is one job per day. total difficulty will be 3.
//
//
//
// Constraints:
//
//
// 1 <= jobDifficulty.length <= 300
// 0 <= jobDifficulty[i] <= 1000
// 1 <= d <= 10
//
//
// Related Topics Array Dynamic Programming 👍 1455 👎 166

package main

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
var dp map[int]map[int]int

func minDifficulty(jobDifficulty []int, d int) int {
	n := len(jobDifficulty)
	dp = map[int]map[int]int{}
	if n < d {
		return -1
	}

	return minDifficultyTraverse(0, d, jobDifficulty)
}

func minDifficultyTraverse(index, d int, diff []int) int {
	if v, ok := dp[index][d]; ok {
		return v
	}
	if d == 1 {
		top := diff[index]
		for i := index; i < len(diff); i++ {
			top = max(top, diff[i])
		}
		return top
	}
	result := math.MaxInt
	top := 0
	for i := index; i <= len(diff)-d; i++ {
		top = max(top, diff[i])
		result = min(result, top+minDifficultyTraverse(i+1, d-1, diff))

	}
	if dp[index] == nil {
		dp[index] = map[int]int{}
	}
	dp[index][d] = result
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
