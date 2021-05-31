package h32

// 思路
//	动态规划，设n为s的长度，ml,mr表示最长匹配区间的左右边界
// 当前s的索引为i，depth为当前括号深度，left、right表示前一个匹配区间的起始与结束位置，左闭右开，ans用于保存
// 程序开始遍历s，初始depth为0
//	如果遇到左括号，depth+=1
//	如果遇到右括号，且depth为0，则跳过
//	如果遇到右括号，且depth大于0，depth-=1，ans[i] = ans[i-1]，如果depth减为0了，ans[i] = max(ans[i], ans[i-1])
// 当遇到左括号时，上述方法可以解决(())的情形，但另外一种情况需要合并区间，比如()()这种形式，所以逻辑如下
//	如果遇到左括号，且depth为0，且right==i，说明当前位置与之前区间连续，depth+=1，right+=1，ans[i] = ans[i-1]

// 2021-05-28
//	动态规划，ans[i]位置之前的最长匹配，最长匹配串的右边界可能在i位置，也可能小于i，left、high用来保存右边界在i位置时已匹配的区间，[left,right]，初始都为0
//	逐步读到一个匹配后，扩展匹配
// 2021-05-31
//	初始状态left和right置为-1，从i==1开始判定，如果i为右括号哦，i-1为左括号，则ans[i]为2，令left=i-1、right=i
//	遍历过程中如果left等于right
//
// 重新梳理一下思路
//	核心思想是首先找到最小匹配()，然后向两侧扩展，(...()...)，直到无法扩展后，再看前方是否有可以连接的区间，比如这样(...()...)(...()...)
//	连接前方匹配区域后，继续看是否是被嵌套的区间，比如(...(...()...)(...()...)...)，如此往复下去，可以处理所有兄弟匹配和父级匹配
//	ans保存到当前位置i的最长匹配
func longestValidParentheses(s string) int {
	n := len(s)
	if n < 2 {
		return 0
	}

	ans := make([]int, n)
	left, right := -1, -1

	for i := 1; i < n; i++ {
		if left == right { // left等于right，表明i-1不在匹配中
			if s[i] == ')' && s[i-1] == '(' {
				left, right = i-1, i
				ans[i] = 2 // 一对儿匹配长度为2
				if left > 0 && ans[left-1] > 0 {
					ans[i] += ans[left-1]
					left = i - ans[i] + 1
				}
			}
		} else {
			if s[i] == '(' { // left不等于right，表明i-1是匹配中的右括号，左括号可能意味着开启新一轮匹配
				left, right = -1, -1
			} else if s[i] == ')' { // left不等于right，表明i-1是匹配中的右括号
				if left > 0 && s[left-1] == '(' { // 当前区域是i-1右括号已匹配区域的外沿
					ans[i] = ans[i-1] + 2
					left, right = left-1, i // 扩展已匹配区域
					if left > 0 && ans[left-1] > 0 {
						ans[i] += ans[left-1]
						left = i - ans[i] + 1
					}
				} else {
					left, right = -1, -1
				}
			}
		}
	}

	var max int
	for i := 0; i < n; i++ {
		if ans[i] > max {
			max = ans[i]
		}
	}

	return max
}
