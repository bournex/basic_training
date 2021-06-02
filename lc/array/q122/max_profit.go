package q122

// 思路
//	要收益首先得是低买高卖，如果要可以多次买卖的话，就是捕捉到所有的上升区间
//	任何上升区间都是由原子的上升区间构成的
func maxProfit(prices []int) int {
	n := len(prices)
	profit := 0
	for i := 1; i < n; i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}
