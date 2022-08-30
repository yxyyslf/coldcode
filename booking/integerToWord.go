package booking

import "strings"

//Convert a non-negative integer num to its English words representation.
//
//
// Example 1:
//
//
//Input: num = 123
//Output: "One Hundred Twenty Three"
//
//
// Example 2:
//
//
//Input: num = 12345
//Output: "Twelve Thousand Three Hundred Forty Five"
//
//
// Example 3:
//
//
//Input: num = 1234567
//Output: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty
//Seven"
//
//
//
// Constraints:
//
//
// 0 <= num <= 2³¹ - 1
//
//
// Related Topics Math String Recursion 👍 2309 👎 5405

//leetcode submit region begin(Prohibit modification and deletion)
var underTwenty = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven",
	"Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
var tens = []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
var chunks = []string{"", "Thousand", "Million", "Billion"}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	index := 0
	result := ""
	for num > 0 {
		if num%1000 != 0 {
			s := ""
			helper(num%1000, &s)
			result = s + chunks[index] + " " + result
		}
		num /= 1000
		index++
	}

	result = strings.TrimSpace(result)
	return result
}

func helper(num int, s *string) {
	if num == 0 {
		return
	}
	if num < 20 {
		*s += underTwenty[num] + " "
		return
	} else if num < 100 {
		*s += tens[num/10] + " "
		helper(num%10, s)
		return
	} else {
		*s += underTwenty[num/100] + " Hundred "
		helper(num%100, s)
		return
	}
}

//leetcode submit region end(Prohibit modification and deletion)
