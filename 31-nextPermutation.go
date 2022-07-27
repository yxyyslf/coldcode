package main

import "math"

func nextPermutation(nums []int)  {
	index := 0
	reverseFlag := 0
	for i:=len(nums)-1;i >= 1;i-- {
		if nums[i-1] < nums[i] {
			index = i-1
			reverseFlag = 1
			break
		}
	}
	if reverseFlag==0 {
		permutationReverse(nums)
		return
	}
	changeIndex := index+1
	minVal := math.MaxInt16
	for i:=len(nums)-1;i > index;i-- {
		if nums[i] > nums[index] && nums[i] - nums[index] < minVal{
			minVal = nums[i] - nums[index]
			changeIndex = i
		}
	}

	nums[index], nums[changeIndex] = nums[changeIndex], nums[index]
	permutationReverse(nums[index+1:])
}

func permutationReverse(nums []int) {
	l := 0
	r := len(nums)-1
	for l <= r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}