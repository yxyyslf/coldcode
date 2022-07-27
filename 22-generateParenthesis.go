package main


func generateParenthesis(n int) []string {
	result := make([]string, 0)
	generateParenthesisTraverse(&result, "",n,n)
	return result
}

func generateParenthesisTraverse(result *[]string, s string, left, right int) {
	if left < 0 || right <0 {
		return
	}
	if left == 0 && right == 0 {
		*result = append(*result, s)
		s = ""
		return
	}
	
	s = s + "("
	generateParenthesisTraverse(result, s, left-1, right)
	s = s[:len(s)-1]
	if right > left {
		s = s + ")"
		generateParenthesisTraverse(result, s, left, right-1)
		s = s[:len(s)-1]
	}

}
