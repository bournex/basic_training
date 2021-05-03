package m91

func numDecodings(s string) int {
	n := len(s)
	if n == 0 || s[0] == '0' {
		return 0
	}
	sum := 0
	var flag bool

	var backtrace func(int)
	backtrace = func(idx int) {
		if idx < n && s[idx] == '0' {
			return
		}

		if idx == n {
			sum++
			return
		}

		if idx+1 < n && (s[idx]-'0')*10+s[idx+1]-'0' <= 26 {
			backtrace(idx + 2)
		}

		backtrace(idx + 1)
	}
	backtrace(0)
	if flag {
		return 0
	}
	return sum
}
