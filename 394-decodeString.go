//Given an encoded string, return its decoded string.
//
// The encoding rule is: k[encoded_string], where the encoded_string inside the
//square brackets is being repeated exactly k times. Note that k is guaranteed to
//be a positive integer.
//
// You may assume that the input string is always valid; there are no extra
//white spaces, square brackets are well-formed, etc. Furthermore, you may assume
//that the original data does not contain any digits and that digits are only for
//those repeat numbers, k. For example, there will not be input like 3a or 2[4].
//
// The test cases are generated so that the length of the output will never
//exceed 10⁵.
//
//
// Example 1:
//
//
//Input: s = "3[a]2[bc]"
//Output: "aaabcbc"
//
//
// Example 2:
//
//
//Input: s = "3[a2[c]]"
//Output: "accaccacc"
//
//
// Example 3:
//
//
//Input: s = "2[abc]3[cd]ef"
//Output: "abcabccdcdcdef"
//
//
//
// Constraints:
//
//
// 1 <= s.length <= 30
// s consists of lowercase English letters, digits, and square brackets '[]'.
// s is guaranteed to be a valid input.
// All the integers in s are in the range [1, 300].
//
//
// Related Topics String Stack Recursion 👍 8849 👎 381

package main

import "strings"

//leetcode submit region begin(Prohibit modification and deletion)
type DecodeStruct struct {
	Times   int
	Strings string
}

func decodeString(s string) string {
	stack := make([]DecodeStruct, 0)
	var times int
	var result string
	numMap := map[byte]int{'1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, '0': 0}
	for i := 0; i < len(s); i++ {
		b := s[i]
		if b >= '0' && b <= '9' {
			times = times*10 + numMap[b]
		} else if b == '[' {
			decodeStruct := DecodeStruct{
				Times:   times,
				Strings: result,
			}
			stack = append(stack, decodeStruct)
			times = 0
			result = ""
		} else if b == ']' {
			decodeStruct := stack[len(stack)-1]
			result = decodeStruct.Strings + strings.Repeat(result, decodeStruct.Times)
			stack = stack[:len(stack)-1]
		} else {
			result += string(b)
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
