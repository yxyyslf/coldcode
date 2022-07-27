package main

import "fmt"

func longestValidParentheses(s string) int {
	stack := make([]int, 0)
	flagMap := make([]int, len(s))

	for i:=0;i < len(s);i++ {
		if s[i] == ')' {
			if len(stack) == 0 || s[stack[len(stack)-1]] != '(' {
				flagMap[i] = 1
			} else {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, i)
		}
	}

	for len(stack) > 0 {
		index := stack[len(stack)-1]
		flagMap[index] = 1
		stack = stack[:len(stack)-1]
	}

	maxVal := 0
	count := 0
	fmt.Println(flagMap)
	for i:=0;i < len(flagMap);i++ {
		if flagMap[i] == 0 {
			count++
		} else {
			maxVal = max(maxVal, count)
			count = 0
		}
	}
	return max(count, maxVal)
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}
