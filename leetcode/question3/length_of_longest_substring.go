package question3

func lengthOfLongestSubstring(s string) int {
	sign := make([]int, 256)
	var (
		head int // window head
		tail int // window tail
		max  int = 1
	)

	n := len(s)
	if n <= 1 {
		return n
	}

	for tail < len(s) {
		if sign[s[tail]] != 0 {
			if tail-head > max {
				max = tail - head
			}
			for s[head] != s[tail] {
				sign[s[head]] = 0
				head++
			}
			head++ // 跳过重复的字符
		}
		sign[s[tail]] = tail + 1 // 更新重复字符索引

		tail++
	}
	if tail-head > max {
		max = tail - head
	}

	return max
}
