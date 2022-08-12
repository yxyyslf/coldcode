//Given n non-negative integers representing an elevation map where the width
//of each bar is 1, compute how much water it can trap after raining.
//
//
// Example 1:
//
//
//Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
//Output: 6
//Explanation: The above elevation map (black section) is represented by array [
//0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section)
//are being trapped.
//
//
// Example 2:
//
//
//Input: height = [4,2,0,3,2,5]
//Output: 9
//
//
//
// Constraints:
//
//
// n == height.length
// 1 <= n <= 2 * 10⁴
// 0 <= height[i] <= 10⁵
//
//
// Related Topics Array Two Pointers Dynamic Programming Stack Monotonic Stack ?
//? 20939 👎 291

package main

//leetcode submit region begin(Prohibit modification and deletion)
func trap(height []int) int {
	result := 0
	if height == nil && len(height) < 3 {
		return result
	}
	n := len(height)
	leftMax := height[0]
	rightMax := height[n-1]
	left := 1
	right := n - 2

	for left <= right {
		if rightMax > leftMax {
			if height[left] < leftMax {
				result += leftMax - height[left]
			} else {
				leftMax = height[left]
			}
			left++
		} else {
			if height[right] < rightMax {
				result += rightMax - height[right]
			} else {
				rightMax = height[right]
			}
			right--
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
