package main

/*
func main() {
	coins := []int{1, 2, 5}
	result := coinChange(coins, 11)
	fmt.Printf("Hello and welcome, %d!", result)
	return
}
*/

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i, _ := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for idx := 1; idx <= amount; idx++ {
		for _, choice := range coins {
			if idx-choice < 0 {
				continue
			}
			// 状态转移，求解少一步的子过程
			dp[idx] = min(dp[idx], dp[idx-choice]+1)
		}
	}
	if dp[amount] == amount+1 {
		return -1
	} else {
		return dp[amount]
	}
}
