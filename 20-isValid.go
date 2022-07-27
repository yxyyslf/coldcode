package main

func isValid(s string) bool {
	stack := make([]string, 0)
	for i:=0;i < len(s);i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, string(s[i]))
		} else {
	     	if len(stack) <= 0 {
	     		return false
			}
	     	temp := stack[len(stack)-1]
			if temp == "(" && s[i] != ')' {
				return false
			}
	     	if temp == "{" && s[i] != '}' {
	     		return false
			}
	     	if temp == "[" && s[i] != ']' {
	     		return false
			}
	     	stack = stack[:len(stack)-1]
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}
