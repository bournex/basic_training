package q125

func isPalindrome(s string) bool {
	n := len(s)
	if n <= 1 {
		return true
	}

	lower_letter := func(b byte) byte {
		if b >= 'A' && b <= 'Z' {
			return b + 32
		} else if (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9') {
			return b
		}
		return 0
	}

	i, j := 0, n-1
	for i < j {
		li, lj := lower_letter(s[i]), lower_letter(s[j])
		if li == 0 {
			i++
		} else if lj == 0 {
			j--
		} else {
			if li == lj {
				i++
				j--
			} else {
				return false
			}
		}
	}
	return true
}
