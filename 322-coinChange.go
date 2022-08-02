//You are given an integer array coins representing coins of different
//denominations and an integer amount representing a total amount of money.
//
// Return the fewest number of coins that you need to make up that amount. If
//that amount of money cannot be made up by any combination of the coins, return -1.
//
//
// You may assume that you have an infinite number of each kind of coin.
//
//
// Example 1:
//
//
//Input: coins = [1,2,5], amount = 11
//Output: 3
//Explanation: 11 = 5 + 5 + 1
//
//
// Example 2:
//
//
//Input: coins = [2], amount = 3
//Output: -1
//
//
// Example 3:
//
//
//Input: coins = [1], amount = 0
//Output: 0
//
//
//
// Constraints:
//
//
// 1 <= coins.length <= 12
// 1 <= coins[i] <= 2³¹ - 1
// 0 <= amount <= 10⁴
//
//
// Related Topics Array Dynamic Programming Breadth-First Search 👍 13082 👎 293
//

package main

//leetcode submit region begin(Prohibit modification and deletion)
func coinChange(coins []int, amount int) int {
	coinDp := make([]int, amount+1)
	for i := 0; i < len(coinDp); i++ {
		coinDp[i] = amount + 1
	}
	coinDp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] {
				coinDp[i] = min(coinDp[i], 1+coinDp[i-coins[j]])
			} else {
				continue
			}
		}
	}
	if coinDp[amount] == amount+1 {
		return -1
	}

	return coinDp[amount]
}

//leetcode submit region end(Prohibit modification and deletion)
