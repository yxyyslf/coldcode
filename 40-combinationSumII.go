//Given a collection of candidate numbers (candidates) and a target number (
//target), find all unique combinations in candidates where the candidate numbers sum
//to target.
//
// Each number in candidates may only be used once in the combination.
//
// Note: The solution set must not contain duplicate combinations.
//
//
// Example 1:
//
//
//Input: candidates = [10,1,2,7,6,1,5], target = 8
//Output:
//[
//[1,1,6],
//[1,2,5],
//[1,7],
//[2,6]
//]
//
//
// Example 2:
//
//
//Input: candidates = [2,5,2,1,2], target = 5
//Output:
//[
//[1,2,2],
//[5]
//]
//
//
//
// Constraints:
//
//
// 1 <= candidates.length <= 100
// 1 <= candidates[i] <= 50
// 1 <= target <= 30
//
//
// Related Topics Array Backtracking 👍 6598 👎 163

package main

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func combinationSum2(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	combinationSum2Traverse(0, target, candidates, &path, &result)
	return result
}

func combinationSum2Traverse(index, target int, candidates []int, path *[]int, result *[][]int) {
	if target == 0 {
		temp := make([]int, len(*path))
		copy(temp, *path)
		*result = append(*result, temp)
		return
	}

	for i := index; i < len(candidates); i++ {
		if target-candidates[i] < 0 {
			break
		}

		if i > index && candidates[i] == candidates[i-1] {
			continue
		}
		*path = append(*path, candidates[i])
		combinationSum2Traverse(i+1, target-candidates[i], candidates, path, result)
		*path = (*path)[:len(*path)-1]
	}
}

//leetcode submit region end(Prohibit modification and deletion)
