//Given a string s, partition s such that every substring of the partition is a
//palindrome. Return all possible palindrome partitioning of s.
//
// A palindrome string is a string that reads the same backward as forward.
//
//
// Example 1:
// Input: s = "aab"
//Output: [["a","a","b"],["aa","b"]]
//
// Example 2:
// Input: s = "a"
//Output: [["a"]]
//
//
// Constraints:
//
//
// 1 <= s.length <= 16
// s contains only lowercase English letters.
//
//
// Related Topics String Dynamic Programming Backtracking 👍 7836 👎 238

package main

//leetcode submit region begin(Prohibit modification and deletion)
func partition(s string) [][]string {
	result := make([][]string, 0)
	path := make([]string, 0)
	partitionTraverse(0, s, &path, &result)
	return result
}

func partitionTraverse(index int, s string, path *[]string, result *[][]string) {
	if index == len(s) {
		tmp := make([]string, len(*path))
		copy(tmp, *path)
		*result = append(*result, tmp)
		return
	}
	for i := index; i < len(s); i++ {
		if partitionCheckPalindrome(s, index, i) {
			*path = append(*path, s[index:i+1])
			partitionTraverse(i+1, s, path, result)
			*path = (*path)[:len(*path)-1]
		}
	}

}

func partitionCheckPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
