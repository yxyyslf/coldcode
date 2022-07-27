//Given a string s, return the number of palindromic substrings in it.
//
// A string is a palindrome when it reads the same backward as forward.
//
// A substring is a contiguous sequence of characters within the string.
//
//
// Example 1:
//
//
//Input: s = "abc"
//Output: 3
//Explanation: Three palindromic strings: "a", "b", "c".
//
//
// Example 2:
//
//
//Input: s = "aaa"
//Output: 6
//Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".
//
//
//
// Constraints:
//
//
// 1 <= s.length <= 1000
// s consists of lowercase English letters.
//
//
// Related Topics String Dynamic Programming 👍 7597 👎 167

package main

//leetcode submit region begin(Prohibit modification and deletion)
func countSubstrings(s string) int {
	result := 0
	if len(s) == 0 {
		return result
	}
	for i := range s {
		result += expand(i, i, s)
		result += expand(i, i+1, s)
	}
	return result

}

func expand(left, right int, s string) int {
	count := 0
	for left >= 0 && right < len(s) && s[left] == s[right] {
		count++
		left--
		right++
	}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)
