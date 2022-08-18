//Given a string s and a dictionary of strings wordDict, return true if s can
//be segmented into a space-separated sequence of one or more dictionary words.
//
// Note that the same word in the dictionary may be reused multiple times in
//the segmentation.
//
//
// Example 1:
//
//
//Input: s = "leetcode", wordDict = ["leet","code"]
//Output: true
//Explanation: Return true because "leetcode" can be segmented as "leet code".
//
//
// Example 2:
//
//
//Input: s = "applepenapple", wordDict = ["apple","pen"]
//Output: true
//Explanation: Return true because "applepenapple" can be segmented as "apple
//pen apple".
//Note that you are allowed to reuse a dictionary word.
//
//
// Example 3:
//
//
//Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
//Output: false
//
//
//
// Constraints:
//
//
// 1 <= s.length <= 300
// 1 <= wordDict.length <= 1000
// 1 <= wordDict[i].length <= 20
// s and wordDict[i] consist of only lowercase English letters.
// All the strings of wordDict are unique.
//
//
// Related Topics Hash Table String Dynamic Programming Trie Memoization 👍 1168
//0 👎 504

package main

//leetcode submit region begin(Prohibit modification and deletion)
func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]bool)
	for _, word := range wordDict {
		wordMap[word] = true
	}
	wordBreakDp := make([]bool, len(s)+1)
	wordBreakDp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			_, ok := wordMap[s[j:i]];
			if ok && wordBreakDp[j] {
				wordBreakDp[i] = true
				break
			}
		}
	}

	return wordBreakDp[len(s)]
}

//leetcode submit region end(Prohibit modification and deletion)
