package solution

import "sort"

//给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
//
//计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
//
//你可以认为每种硬币的数量是无限的。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/coin-change
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//超时
//本题使用记忆搜索或者动态规划解题，可以暂时不做
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if len(coins) == 1 {
		if amount%coins[0] == 0 {
			return amount / coins[0]
		}
		return -1
	}
	sort.Ints(coins)
	retNum := -1
	maxNum := amount / coins[len(coins)-1]
	for i := maxNum; i >= 0; i-- {
		remainder := amount - coins[len(coins)-1]*i
		coinNum := coinChange(coins[:len(coins)-1], remainder)
		if coinNum == -1 {
			continue
		}
		if coinNum+i < retNum || retNum == -1 {
			retNum = coinNum + i
		}
	}
	return retNum
}
