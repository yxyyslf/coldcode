//Given a list of unique words, return all the pairs of the distinct indices (i,
// j) in the given list, so that the concatenation of the two words words[i] +
//words[j] is a palindrome.
//
//
// Example 1:
//
//
//Input: words = ["abcd","dcba","lls","s","sssll"]
//Output: [[0,1],[1,0],[3,2],[2,4]]
//Explanation: The palindromes are ["dcbaabcd","abcddcba","slls","llssssll"]
//
//
// Example 2:
//
//
//Input: words = ["bat","tab","cat"]
//Output: [[0,1],[1,0]]
//Explanation: The palindromes are ["battab","tabbat"]
//
//
// Example 3:
//
//
//Input: words = ["a",""]
//Output: [[0,1],[1,0]]
//
//
//
// Constraints:
//
//
// 1 <= words.length <= 5000
// 0 <= words[i].length <= 300
// words[i] consists of lower-case English letters.
//
//
// Related Topics Array Hash Table String Trie 👍 2948 👎 286

package main

//leetcode submit region begin(Prohibit modification and deletion)
func palindromePairs(words []string) [][]int {
	hash := map[string]int{}
	result := [][]int{}

	for i, word := range words {
		hash[word] = i
	}

	for i, word := range words {
		// 字符前半部分回文 后半部分hash存在
		for j := 0; j < len(word); j++ {
			if v, ok := hash[pairReverse(word[j+1:])]; ok && isPairPalindrome(word, 0, j) && i != v {
				result = append(result, []int{v, i})
				if pairReverse(word[j+1:]) == "" {
					result = append(result, []int{i, v})
				}
			}
			if v, ok := hash[pairReverse(word[:j+1])]; ok && isPairPalindrome(word, j+1, len(word)-1) && i != v {
				result = append(result, []int{i, v})
			}
		}
	}

	return result
}

func isPairPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func pairReverse(s string) string {
	chars := []byte(s)
	left, right := 0, len(s)-1
	for left < right {
		chars[left], chars[right] = chars[right], chars[left]
		left++
		right--
	}
	return string(chars)
}

//leetcode submit region end(Prohibit modification and deletion)
