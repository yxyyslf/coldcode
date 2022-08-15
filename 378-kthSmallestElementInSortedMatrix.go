//Given an n x n matrix where each of the rows and columns is sorted in
//ascending order, return the kᵗʰ smallest element in the matrix.
//
// Note that it is the kᵗʰ smallest element in the sorted order, not the kᵗʰ
//distinct element.
//
// You must find a solution with a memory complexity better than O(n²).
//
//
// Example 1:
//
//
//Input: matrix = [[1,5,9],[10,11,13],[12,13,15]], k = 8
//Output: 13
//Explanation: The elements in the matrix are [1,5,9,10,11,12,13,13,15], and
//the 8ᵗʰ smallest number is 13
//
//
// Example 2:
//
//
//Input: matrix = [[-5]], k = 1
//Output: -5
//
//
//
// Constraints:
//
//
// n == matrix.length == matrix[i].length
// 1 <= n <= 300
// -10⁹ <= matrix[i][j] <= 10⁹
// All the rows and columns of matrix are guaranteed to be sorted in non-
//decreasing order.
// 1 <= k <= n²
//
//
//
// Follow up:
//
//
// Could you solve the problem with a constant memory (i.e., O(1) memory
//complexity)?
// Could you solve the problem in O(n) time complexity? The solution may be too
//advanced for an interview but you may find reading this paper fun.
//
//
// Related Topics Array Binary Search Sorting Heap (Priority Queue) Matrix 👍 79
//49 👎 287

package main

//leetcode submit region begin(Prohibit modification and deletion)
func kthSmallestInSortedMatrix(matrix [][]int, k int) int {
	m := len(matrix)
	n := len(matrix[0])
	left := matrix[0][0]
	right := matrix[m-1][n-1]
	for left < right {
		mid := left + (right-left)/2
		count := checkSmallNumber(matrix, mid, m, n)
		if count < k {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func checkSmallNumber(matrix [][]int, mid, m, n int) int {
	i, j, count := m-1, 0, 0

	for i >= 0 && j < n {
		if matrix[i][j] <= mid {
			count += i + 1
			j++
		} else {
			i--
		}
	}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)
