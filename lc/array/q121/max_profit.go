package q121

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
// 思路
//	要想盈利，需要卖出价格比买入价格高，如果要收益最大化，就需要在所有的低买高卖里，价差最高
//	暴力解法下，遍历遍历前n-1天作为买入点，在每一个买入点之后，遍历所有可以卖出的价格，找出最高收益的买卖时机
//	暴力解法复杂度是 N²/2，空间复杂度为常数
//	观察发现在内循环中存在重复扫描，可以保存第i天后的最高价格，避免每次扫描，两趟遍历解决问题，时间复杂度为N，空间复杂度为N，空间换时间
//	上面的方法是通过反向先找到i位置之后的最大值，需要两遍遍历和N的空间复杂度，如果从前向后扫描，改为仅保存当前的极小值，则只需一次遍历，且常数空间

func maxProfit(prices []int) int {
	n := len(prices)
	// min := prices[0]
	max := 0

	for i := 1; i < n; i++ {
		if prices[i] < prices[0] {
			prices[0] = prices[i]
		} else if prices[i] > max+prices[0] {
			max = prices[i] - prices[0]
		}
	}
	return max

	// n := len(prices)
	// min := prices[0]
	// max := 0

	// for i := 1; i < n; i++ {
	// 	if prices[i] < min {
	// 		min = prices[i]
	// 	} else if prices[i] > max+min {
	// 		max = prices[i] - min
	// 	}
	// }
	// return max
}

// 空间换时间解法
func maxProfit1(prices []int) int {
	n := len(prices)
	maxprice := make([]int, n)

	// 获取i位置之后的最大值数组
	maxprice[n-1] = prices[n-1]
	for i := n - 2; i >= 0; i-- {
		if prices[i] > maxprice[i+1] {
			maxprice[i] = prices[i]
		} else {
			maxprice[i] = maxprice[i+1]
		}
	}

	var max int
	for i := 0; i < n-1; i++ {
		if maxprice[i]-prices[i] > max {
			max = maxprice[i] - prices[i]
		}
	}

	return max
}
