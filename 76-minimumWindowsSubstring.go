//Given two strings s and t of lengths m and n respectively, return the minimum
//window substring of s such that every character in t (including duplicates) is
//included in the window. If there is no such substring, return the empty string
//"".
//
// The testcases will be generated such that the answer is unique.
//
// A substring is a contiguous sequence of characters within the string.
//
//
// Example 1:
//
//
//Input: s = "ADOBECODEBANC", t = "ABC"
//Output: "BANC"
//Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C'
//from string t.
//
//
// Example 2:
//
//
//Input: s = "a", t = "a"
//Output: "a"
//Explanation: The entire string s is the minimum window.
//
//
// Example 3:
//
//
//Input: s = "a", t = "aa"
//Output: ""
//Explanation: Both 'a's from t must be included in the window.
//Since the largest window of s only has one 'a', return empty string.
//
//
//
// Constraints:
//
//
// m == s.length
// n == t.length
// 1 <= m, n <= 10⁵
// s and t consist of uppercase and lowercase English letters.
//
//
//
//Follow up: Could you find an algorithm that runs in
//O(m + n) time?
//
// Related Topics Hash Table String Sliding Window 👍 11631 👎 555

package main

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
func minWindow(s string, t string) string {
	minVal := math.MaxInt
	if len(t) > len(s) {
		return ""
	}
	var left, right, count, start int
	letterMap := make(map[byte]int, 0)
	windowMap := make(map[byte]int, 0)
	for i := range t {
		letterMap[t[i]]++
	}
	for right < len(s) {
		c := s[right]
		right++
		if letterMap[c] > 0 {
			windowMap[c]++
			if letterMap[c] == windowMap[c] {
				count++
			}
		}

		for count == len(letterMap) {
			if minVal > right-left {
				start = left
				minVal = right - left
			}

			d := s[left]
			left++
			if letterMap[d] > 0 {
				if letterMap[d] == windowMap[d] {
					count--
				}
				windowMap[d]--

			}
		}
	}
	if minVal == math.MaxInt {
		return ""
	}
	return s[start : start+minVal]
}

//leetcode submit region end(Prohibit modification and deletion)
