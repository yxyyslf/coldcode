//Given a string containing digits from 2-9 inclusive, return all possible
//letter combinations that the number could represent. Return the answer in any order.
//
//
// A mapping of digits to letters (just like on the telephone buttons) is given
//below. Note that 1 does not map to any letters.
//
//
// Example 1:
//
//
//Input: digits = "23"
//Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
//
//
// Example 2:
//
//
//Input: digits = ""
//Output: []
//
//
// Example 3:
//
//
//Input: digits = "2"
//Output: ["a","b","c"]
//
//
//
// Constraints:
//
//
// 0 <= digits.length <= 4
// digits[i] is a digit in the range ['2', '9'].
//
//
// Related Topics Hash Table String Backtracking 👍 12002 👎 739

package main

//leetcode submit region begin(Prohibit modification and deletion)
func letterCombinations(digits string) []string {

	result := make([]string, 0)
	if len(digits) == 0 {
		return result
	}
	letterMap := map[byte]string{
		'1': "", '2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz", '0': "",
	}
	letterCombinationsTraverse(0, "", digits, letterMap, &result)
	return result
}

func letterCombinationsTraverse(index int, str, digits string, m map[byte]string, result *[]string) {
	if index == len(digits) {
		*result = append(*result, str)
		return
	}
	letters := m[digits[index]]
	for _, letter := range letters {
		str += string(letter)
		letterCombinationsTraverse(index+1, str, digits, m, result)
		str = str[:len(str)-1]
	}
}

//leetcode submit region end(Prohibit modification and deletion)
