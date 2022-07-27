//You are given a string s. We want to partition the string into as many parts
//as possible so that each letter appears in at most one part.
//
// Note that the partition is done so that after concatenating all the parts in
//order, the resultant string should be s.
//
// Return a list of integers representing the size of these parts.
//
//
// Example 1:
//
//
//Input: s = "ababcbacadefegdehijhklij"
//Output: [9,7,8]
//Explanation:
//The partition is "ababcbaca", "defegde", "hijhklij".
//This is a partition so that each letter appears in at most one part.
//A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it
//splits s into less parts.
//
//
// Example 2:
//
//
//Input: s = "eccbbbbdec"
//Output: [10]
//
//
//
// Constraints:
//
//
// 1 <= s.length <= 500
// s consists of lowercase English letters.
//
//
// Related Topics Hash Table Two Pointers String Greedy 👍 8093 👎 309

package main

//leetcode submit region begin(Prohibit modification and deletion)
func partitionLabels(s string) []int {
	result := make([]int, 0)
	if len(s) == 0 {
		return result
	}

	letter := [26]int{}
	for i := range s {
		letter[s[i]-'a'] = i
	}
	left := 0
	right := 0
	for i := range s {
		if right < letter[s[i]-'a'] {
			right = letter[s[i]-'a']
		}
		if i == right {
			result = append(result, right-left+1)
			left = i + 1
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
