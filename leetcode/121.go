package leetcode

/**
Input: [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
             Not 7-1 = 6, as selling price needs to be larger than buying price.
Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
 */

func maxProfit(prices []int) int {
	max_num := 0
	for i := 0; i < len(prices); i++ {
		for j := i+1; j < len(prices); j++ {
			if max_num < (prices[j] - prices[i]) {
				max_num = prices[j] - prices[i]
			}
		}
	}
	return max_num
}