//Given two strings s and p, return an array of all the start indices of p's
//anagrams in s. You may return the answer in any order.
//
// An Anagram is a word or phrase formed by rearranging the letters of a
//different word or phrase, typically using all the original letters exactly once.
//
//
// Example 1:
//
//
//Input: s = "cbaebabacd", p = "abc"
//Output: [0,6]
//Explanation:
//The substring with start index = 0 is "cba", which is an anagram of "abc".
//The substring with start index = 6 is "bac", which is an anagram of "abc".
//
//
// Example 2:
//
//
//Input: s = "abab", p = "ab"
//Output: [0,1,2]
//Explanation:
//The substring with start index = 0 is "ab", which is an anagram of "ab".
//The substring with start index = 1 is "ba", which is an anagram of "ab".
//The substring with start index = 2 is "ab", which is an anagram of "ab".
//
//
//
// Constraints:
//
//
// 1 <= s.length, p.length <= 3 * 10⁴
// s and p consist of lowercase English letters.
//
//
// Related Topics Hash Table String Sliding Window 👍 8189 👎 266

package main

//leetcode submit region begin(Prohibit modification and deletion)
func findAnagrams(s string, p string) []int {
	result := make([]int, 0)
	letterMap := [26]int{}
	for i := 0; i < len(p); i++ {
		letterMap[p[i]-'a']++
	}

	for i := 0; i <= len(s)-len(p); i++ {
		flag := true
		subMap := [26]int{}
		sub := s[i : i+len(p)]
		for j := 0; j < len(sub); j++ {
			subMap[sub[j]-'a']++
		}
		for k := 0; k < 26; k++ {
			if subMap[k] != letterMap[k] {
				flag = false
				break
			}
		}
		if flag {
			result = append(result, i)
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
