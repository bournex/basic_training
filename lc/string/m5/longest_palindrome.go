package m5

// 思考
//	看过了官方答案，选择中心扩散来实现，因为自己能思考到的唯一方法，就是中心扩散
//	所以基本上就是边界处理的问题

func longestPalindrome(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	isPalindrome := func(left, right int) (int, int) {
		for left >= 0 && right <= n-1 && s[left] == s[right] {
			left--
			right++
		}
		return left + 1, right - 1
	}

	l, r, mr, ml := 0, 0, 0, 0
	for i := 1; i < n; i++ {
		l, r = isPalindrome(i, i)
		if r-l > mr-ml {
			mr, ml = r, l
		}
		l, r = isPalindrome(i-1, i)
		if r-l > mr-ml {
			mr, ml = r, l
		}
	}

	if ml == mr {
		return s[0:1]
	}
	return s[ml : mr+1]
}
