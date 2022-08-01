package main

func findTargetSumWays(nums []int, target int) int {
	result = 0
	if nums == nil || len(nums) == 0 {
		return result
	}
	traverseSumWays(0, target, nums)
	return result
}

func traverseSumWays(index, target int, nums []int) {
	if index == len(nums) {
		if target == 0 {
			result++
		}
		return
	}
	traverseSumWays(index+1, target+nums[index], nums)
	traverseSumWays(index+1, target-nums[index], nums)
}
