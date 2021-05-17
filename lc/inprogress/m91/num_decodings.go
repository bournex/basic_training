package m91

func numDecodings(s string) int {
	// 要求得k个字符的解码序列数量n，
	// 当第k个字符单独1个字符解码时，序列数量为k-1时的解码序列数量n1，
	// 当第k个字符与第k-1个字符组合在一起解码时，序列数量为k-2时的解码序列数量n2
	// 如果第k个字符是0，则只能有第二种情形，
	// 如果第k-1个字符是0，则只能有第一种情形，
	// 如果第k、k-1个字符都是0，或者k-1与k构成的两位数大于26，则解码失败，返回0
	n := len(s)
	if n <= 0 || s[0] == '0' {
		return 0
	}

	ans := make([]int, len(s))
	ans[0] = 1

	for i := 1; i < n; i++ {
		if s[i] != '0' {
			ans[i] += ans[i-1]
		}
		if (s[i-1]-'0')*10+(s[i]-'0') < 26 && s[i-1] != '0' {
			if i > 1 {
				ans[i] += ans[i-2]
			} else {
				ans[i]++
			}
		}
		if ans[i] == 0 {
			return 0
		}
	}

	return ans[n-1]
}

// 思考
//	以下是第一版实现，用的是回溯，小规模数据没有问题，大规模数据会超时
//	忽然想到在小破站看的视频，求数量的情况，可能是动态规划的一种题型，所以放弃回溯法
func numDecodings1(s string) int {
	n := len(s)
	if n == 0 || s[0] == '0' {
		return 0
	}
	sum := 0

	var backtrace func(int) bool
	backtrace = func(idx int) bool {
		if idx == n {
			sum++
		}

		var exist bool
		if idx < n {
			if '0' < s[idx] && s[idx] <= '9' {
				exist = exist || backtrace(idx+1)
			}
		}
		if idx+1 < n {
			if s[idx] != '0' && (s[idx]-'0')*10+s[idx+1]-'0' <= 26 {
				exist = exist || backtrace(idx+2)
			}
		}

		return exist
	}

	exist := backtrace(0)
	if exist {
		return 0
	}
	return sum
}
