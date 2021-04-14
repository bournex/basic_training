package q67

func addBinary(a string, b string) string {
	la := len(a)
	lb := len(b)
	if la == 0 || (la == 1 && a[0] == '0') {
		return b
	}
	if lb == 0 || (lb == 1 && b[0] == '0') {
		return a
	}

	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	result := make([]byte, max(la, lb)+1)
	carrier := byte('0')
	for i := max(la, lb) - 1; i >= 0; i-- {
		ka, kb := byte('0'), byte('0')
		if la > 0 {
			la--
			ka = a[la]
		}
		if lb > 0 {
			lb--
			kb = b[lb]
		}
		switch ka + kb + carrier {
		case '0' + '0' + '0':
			result[i+1] = '0'
		case '0' + '0' + '1':
			result[i+1] = '1'
			carrier = '0'
		case '0' + '1' + '1':
			result[i+1] = '0'
			carrier = '1'
		case '1' + '1' + '1':
			result[i+1] = '1'
			carrier = '1'
		}
	}

	if carrier == '1' {
		result[0] = '1'
	} else {
		return string(result[1:])
	}
	return string(result)
}
